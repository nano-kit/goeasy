package impl

import (
	"context"

	"github.com/micro/go-micro/v2/client"
	"github.com/nano-kit/goeasy/internal/namespace"
)

type clientWrapper struct {
	client.Client
	namespace string
}

func (a *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	if len(a.namespace) > 0 {
		ctx = namespace.ContextWithNamespace(ctx, a.namespace)
	}
	return a.Client.Call(ctx, req, rsp, opts...)
}
