package metric

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/nano-kit/goeasy/gate/metric/comet"
	"github.com/nano-kit/goeasy/gate/metric/longpoll"
	"github.com/nano-kit/goeasy/gate/metric/model"
	"github.com/nano-kit/goeasy/gate/metric/rest"
)

// HTTPWrapper returns an measuring standard http.Handler.
func HTTPWrapper(handler http.Handler) httpWrapper {
	rest := New(Config{
		Recorder: rest.NewRecorder(rest.Config{
			DurationBuckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
		}),
	})
	longpoll := New(Config{
		Recorder: longpoll.NewRecorder(longpoll.Config{
			DurationBuckets: []float64{5, 20, 40, 60, 80, 100, 120, 200, 300, 500, 800},
		}),
	})
	comet := New(Config{
		Recorder: comet.NewRecorder(comet.Config{
			DurationBuckets: []float64{5, 60, 300, 600, 1800, 3600, 7200, 43200, 86400, 172800, 604800},
		}),
	})

	return httpWrapper{
		rest:     rest,
		longpoll: longpoll,
		comet:    comet,
		handler:  handler,
	}
}

type httpWrapper struct {
	rest     Middleware
	longpoll Middleware
	comet    Middleware
	handler  http.Handler
}

func (h httpWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		handlerID string
		m         Middleware
	)

	// the long-polling URL paths have a dedicated handler ID and middle ware
	switch r.URL.Path {
	case "/comet/subscribe":
		handlerID = r.URL.Path
		m = h.comet
	case "/liveroom/room/recv":
		handlerID = r.URL.Path
		m = h.longpoll
	default:
		handlerID = "api"
		m = h.rest
	}

	// starts to measure
	wi := &responseWriterInterceptor{
		statusCode:     http.StatusOK,
		ResponseWriter: w,
	}
	reporter := &stdReporter{
		w: wi,
		r: r,
	}

	m.Measure(handlerID, reporter, func() {
		h.handler.ServeHTTP(wi, r)
	})
}

// Config is the configuration for the middleware factory.
type Config struct {
	// Recorder is the way the metrics will be recorder in the different backends.
	Recorder model.Recorder
	// Service is an optional identifier for the metrics, this can be useful if
	// a same service has multiple servers (e.g API, metrics and healthchecks).
	Service string
	// GroupedStatus will group the status label in the form of `\dxx`, for example,
	// 200, 201, and 203 will have the label `code="2xx"`. This impacts on the cardinality
	// of the metrics and also improves the performance of queries that are grouped by
	// status code because there are already aggregated in the metric.
	// By default will be false.
	GroupedStatus bool
	// DisableMeasureSize will disable the recording metrics about the response size,
	// by default measuring size is enabled (`DisableMeasureSize` is false).
	DisableMeasureSize bool
	// DisableMeasureInflight will disable the recording metrics about the inflight requests number,
	// by default measuring inflights is enabled (`DisableMeasureInflight` is false).
	DisableMeasureInflight bool
}

func (c *Config) defaults() {
	if c.Recorder == nil {
		c.Recorder = model.Dummy
	}
}

// Middleware is a service that knows how to measure an HTTP handler by wrapping
// another handler.
//
// Depending on the framework/library we want to measure, this can change a lot,
// to abstract the way how we measure on the different libraries, Middleware will
// recieve a `Reporter` that knows how to get the data the Middleware service needs
// to measure.
type Middleware struct {
	cfg Config
}

// New returns the a Middleware service.
func New(cfg Config) Middleware {
	cfg.defaults()

	m := Middleware{cfg: cfg}

	return m
}

// Measure abstracts the HTTP handler implementation by only requesting a reporter, this
// reporter will return the required data to be measured.
// it accepts a next function that will be called as the wrapped logic before and after
// measurement actions.
func (m Middleware) Measure(handlerID string, reporter Reporter, next func()) {
	ctx := reporter.Context()

	// If there isn't predefined handler ID we
	// set that ID as the URL path.
	hid := handlerID
	if handlerID == "" {
		hid = reporter.URLPath()
	}

	// Measure inflights if required.
	if !m.cfg.DisableMeasureInflight {
		props := model.HTTPProperties{
			Service: m.cfg.Service,
			ID:      hid,
		}
		m.cfg.Recorder.AddInflightRequests(ctx, props, 1)
		defer m.cfg.Recorder.AddInflightRequests(ctx, props, -1)
	}

	// Start the timer and when finishing measure the duration.
	start := time.Now()
	defer func() {
		duration := time.Since(start)

		// If we need to group the status code, it uses the
		// first number of the status code because is the least
		// required identification way.
		var code string
		if m.cfg.GroupedStatus {
			code = fmt.Sprintf("%dxx", reporter.StatusCode()/100)
		} else {
			code = strconv.Itoa(reporter.StatusCode())
		}

		props := model.HTTPReqProperties{
			Service: m.cfg.Service,
			ID:      hid,
			Method:  reporter.Method(),
			Code:    code,
		}
		m.cfg.Recorder.ObserveHTTPRequestDuration(ctx, props, duration)

		// Measure size of response if required.
		if !m.cfg.DisableMeasureSize {
			m.cfg.Recorder.ObserveHTTPResponseSize(ctx, props, reporter.BytesWritten())
		}
	}()

	// Call the wrapped logic.
	next()
}

// Reporter knows how to report the data to the Middleware so it can measure the
// different framework/libraries.
type Reporter interface {
	Method() string
	Context() context.Context
	URLPath() string
	StatusCode() int
	BytesWritten() int64
}

type stdReporter struct {
	w *responseWriterInterceptor
	r *http.Request
}

func (s *stdReporter) Method() string { return s.r.Method }

func (s *stdReporter) Context() context.Context { return s.r.Context() }

func (s *stdReporter) URLPath() string { return s.r.URL.Path }

func (s *stdReporter) StatusCode() int { return s.w.statusCode }

func (s *stdReporter) BytesWritten() int64 { return int64(s.w.bytesWritten) }

// responseWriterInterceptor is a simple wrapper to intercept set data on a
// ResponseWriter.
type responseWriterInterceptor struct {
	http.ResponseWriter
	statusCode   int
	bytesWritten int
}

func (w *responseWriterInterceptor) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *responseWriterInterceptor) Write(p []byte) (int, error) {
	w.bytesWritten += len(p)
	return w.ResponseWriter.Write(p)
}

func (w *responseWriterInterceptor) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, ok := w.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("type assertion failed http.ResponseWriter not a http.Hijacker")
	}
	return h.Hijack()
}

func (w *responseWriterInterceptor) Flush() {
	f, ok := w.ResponseWriter.(http.Flusher)
	if !ok {
		return
	}

	f.Flush()
}

// Check interface implementations.
var (
	_ http.ResponseWriter = &responseWriterInterceptor{}
	_ http.Hijacker       = &responseWriterInterceptor{}
	_ http.Flusher        = &responseWriterInterceptor{}
)
