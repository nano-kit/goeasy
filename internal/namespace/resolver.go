package namespace

import "net/http"

func NewResolver(srvType, namespace string) *Resolver {
	return &Resolver{srvType, namespace}
}

// Resolver determines the namespace for a request
type Resolver struct {
	srvType   string
	namespace string
}

func (r Resolver) String() string {
	return "internal/namespace"
}

func (r Resolver) ResolveWithType(req *http.Request) string {
	return r.Resolve(req) + "." + r.srvType
}

func (r Resolver) Resolve(req *http.Request) string {
	// check to see what the provided namespace is, we only do
	// domain mapping if the namespace is set to 'domain'
	if r.namespace != "domain" {
		return r.namespace
	}

	panic("the domain resolver is not ready")
}

func (r Resolver) Namespace() string {
	return r.namespace
}
