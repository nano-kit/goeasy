package impl

import (
	"context"

	"github.com/micro/go-micro/v2/logger"
	ijson "github.com/nano-kit/goeasy/internal/json"
	"github.com/nano-kit/goeasy/servers/liveuser"
)

type Wx struct {
}

func (w *Wx) Login(ctx context.Context, req *liveuser.LoginReq, res *liveuser.LoginRes) error {
	logger.Infof("got wx.login request: %v", ijson.Stringify(req))
	return nil
}
