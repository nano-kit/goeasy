package liveroom

import (
	context "context"
	"fmt"

	"github.com/micro/go-micro/v2/auth"
	log "github.com/micro/go-micro/v2/logger"
)

type Demo struct{}

func (d *Demo) Hello(ctx context.Context, req *HelloReq, res *HelloRes) error {
	acc, ok := auth.AccountFromContext(ctx)
	log.Infof("Entering Demo.Hello with req=%v acc=%v ok=%v", req, acc, ok)
	res.Ack = fmt.Sprintf("<ack> %s, <account> %+v", req.Say, acc)
	return nil
}
