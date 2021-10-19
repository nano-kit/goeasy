package api

import (
	"net/http"
	"strings"

	"github.com/micro/go-micro/v2/api/resolver"
)

// Resolver is the default resolver for legacy purposes
// it uses proxy routing to resolve names
// /foo becomes namespace.foo
// /v1/foo becomes namespace.v1.foo
type Resolver struct {
	Options resolver.Options
}

var webHandlerPaths = []struct {
	path   string
	prefix bool
}{
	{path: ""},
	{path: "/"},
	{path: "/favicon.ico"},
	{path: "/metrics"},
	{path: "/placeholder"},
	{path: "/o", prefix: true},
	{path: "/portal", prefix: true},
}

func (r *Resolver) Resolve(req *http.Request) (*resolver.Endpoint, error) {
	// resolve as web handler
	for _, x := range webHandlerPaths {
		if x.path == req.URL.Path {
			return nil, resolver.ErrInvalidPath
		}
		if x.prefix && strings.HasPrefix(req.URL.Path, x.path+"/") {
			return nil, resolver.ErrInvalidPath
		}
	}

	// resolve as api handler
	var name, method string

	switch r.Options.Handler {
	// internal handlers
	case "meta", "api", "rpc", "micro":
		name, method = apiRoute(req.URL.Path)
	default:
		method = req.Method
		name = proxyRoute(req.URL.Path)
	}

	ns := r.Options.Namespace(req)
	return &resolver.Endpoint{
		Name:   ns + "." + name,
		Method: method,
	}, nil
}

func (r *Resolver) String() string {
	return "goeasy"
}

// NewResolver creates a new micro resolver
func NewResolver(opts ...resolver.Option) resolver.Resolver {
	return &Resolver{
		Options: resolver.NewOptions(opts...),
	}
}
