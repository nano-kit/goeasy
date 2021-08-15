package gate

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/api/resolver"
	"github.com/micro/go-micro/v2/api/router"
	"github.com/micro/go-micro/v2/api/router/registry"
	"github.com/micro/go-micro/v2/api/server"
	httpapi "github.com/micro/go-micro/v2/api/server/http"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/nano-kit/goeasy/gate/auth"
	iconf "github.com/nano-kit/goeasy/internal/config"
	"github.com/nano-kit/goeasy/internal/handler"
	"github.com/nano-kit/goeasy/internal/namespace"
	"github.com/nano-kit/goeasy/internal/resolver/api"
	"github.com/nano-kit/goeasy/internal/rlimit"
)

const (
	ServiceName = "gate"
	APIPath     = "/"
)

type Server struct {
	Address        string   `json:"api_serving_address"`
	Namespace      string   `json:"namespace"`
	Production     bool     `json:"production"`
	LogOutputPaths []string `json:"logging_output_paths"`
	Domain         string   `json:"domain"`
	EnableTLS      bool     `json:"enable_api_tls"`
	EnableACME     bool     `json:"enable_api_acme"`
}

func NewServer() *Server {
	s := &Server{}
	if err := iconf.LoadInitialConfigFromFile("serverinit.yaml", s); err != nil {
		panic(err)
	}
	return s
}

func (s *Server) Name() string {
	return s.Namespace + "." + ServiceName
}

func (s *Server) Run() {
	log.Init(
		log.WithFields(map[string]interface{}{"service": ServiceName}),
		log.SetOption("outputs", s.LogOutputPaths),
		log.SetOption("color", !s.Production),
	)

	// set max open files
	_ = rlimit.SetNumFiles(10240)

	// initialize the micro service
	var srvOpts []micro.Option
	srvOpts = append(srvOpts, micro.Name(s.Name()))
	service := micro.NewService(srvOpts...)

	// initialize the API gate server
	var opts []server.Option
	opts = append(opts,
		server.EnableCORS(true),
		server.KeepaliveTimeout(10*time.Second),
	)
	opts = s.configureTLS(opts)

	// create the router
	var httpHandler http.Handler
	muxRouter := mux.NewRouter()
	httpHandler = muxRouter

	// return version at root path
	muxRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			return
		}

		response := fmt.Sprintf(`{"version": "%s"}`, "TODO")
		w.Write([]byte(response))
	})
	// strip favicon.ico
	muxRouter.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})

	// create the namespace resolver
	nsResolver := namespace.NewResolver("service", s.Namespace)
	// resolver options
	ropts := []resolver.Option{
		resolver.WithNamespace(nsResolver.ResolveWithType),
		resolver.WithHandler("meta"),
	}
	// default resolver
	rr := api.NewResolver(ropts...)
	log.Infof("Registering API Default Handler at %s", APIPath)
	rt := registry.NewRouter(
		router.WithResolver(rr),
		router.WithRegistry(service.Options().Registry),
	)
	muxRouter.PathPrefix(APIPath).Handler(handler.Meta(service, rt, nsResolver.ResolveWithType))

	// create the auth wrapper and the server
	authWrapper := auth.Wrapper(rr, nsResolver)
	gate := httpapi.NewServer(s.Address, server.WrapHandler(authWrapper))
	gate.Init(opts...)
	gate.Handle("/", httpHandler)

	// Start API Gate
	if err := gate.Start(); err != nil {
		log.Fatal(err)
	}

	// Run micro server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

	// Stop API
	if err := gate.Stop(); err != nil {
		log.Fatal(err)
	}
}
