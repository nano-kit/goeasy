package comet

import (
	"context"
	"fmt"
	"time"

	"github.com/micro/go-micro/v2/auth"
	"github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/util/pubsub"
	iauth "github.com/nano-kit/goeasy/internal/auth"
	"github.com/nano-kit/goeasy/internal/rmgr"
)

// We must hear subscriber's heartbeat within this duration
const heartbeatDuration = 1 * time.Minute

type Comet struct {
	g  *pubsub.Group
	rm *rmgr.Bucket
}

func NewComet() *Comet {
	return &Comet{
		g:  pubsub.New(),
		rm: rmgr.NewBucket(),
	}
}

type streamCtx struct {
	account   *auth.Account
	stream    Comet_SubscribeStream
	cancel    context.CancelFunc
	heartbeat *time.Time
}

type streamCtxKey struct{}

type serverPush struct {
	*PublishReq
}

func (m *serverPush) Topic() string {
	return m.Uid
}

func (m *serverPush) Body() interface{} {
	return m.Evt
}

func (c *Comet) Publish(ctx context.Context, req *PublishReq, res *PublishRes) error {
	c.g.Publish(ctx, &serverPush{req})
	return nil
}

func (c *Comet) Subscribe(ctx context.Context, stream Comet_SubscribeStream) error {
	req, err := stream.Recv()
	if err != nil {
		return errors.BadRequest("incorrect-protocol", "stream recv: %v", err)
	}
	if req.T != MsgType_AUTH {
		return errors.BadRequest("incorrect-protocol", "expect message type AUTH: %v", req)
	}
	account, ok := iauth.AccountFromToken(req.GetAuth().GetToken())
	if !ok {
		return errors.BadRequest("unidentified-subscriber", "")
	}
	logger.Infof("subscriber %q enter", account.ID)

	heartbeat := time.Now()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	ctx = context.WithValue(ctx, streamCtxKey{}, streamCtx{account, stream, cancel, &heartbeat})
	c.g.Go(ctx, c.recv)
	c.g.Subscribe(ctx, account.ID, c.send, pubsub.WithTicker(heartbeatDuration, c.tick))
	return nil
}

func (c *Comet) recv(ctx context.Context) (err error) {
	sc := ctx.Value(streamCtxKey{}).(streamCtx)
	defer func() { sc.cancel() }()
	for {
		var uplink *Uplink
		if uplink, err = sc.stream.Recv(); err != nil {
			return fmt.Errorf("process %q uplink message: stream recv: %v", sc.account.ID, err)
		}
		logger.Debugf("RX %q %v", sc.account.ID, uplink)
		*sc.heartbeat = time.Now()
		switch uplink.T {
		case MsgType_HB:
		case MsgType_JOIN:
		}
	}
}

func (c *Comet) tick(ctx context.Context) (err error) {
	sc := ctx.Value(streamCtxKey{}).(streamCtx)
	defer func() {
		if err != nil {
			sc.cancel()
		}
	}()
	if sc.heartbeat == nil {
		return fmt.Errorf("server tick %q: no heartbeat", sc.account.ID)
	}
	if time.Since(*sc.heartbeat) > 2*heartbeatDuration {
		return fmt.Errorf("server tick %q: heartbeat delays", sc.account.ID)
	}
	if err := sc.stream.Send(&Downlink{
		T:  MsgType_HB,
		Hb: &Heartbeat{},
	}); err != nil {
		return fmt.Errorf("server tick %q: send heartbeat: %v", sc.account.ID, err)
	}
	return nil
}

func (c *Comet) send(ctx context.Context, msg pubsub.Message) (bool, error) {
	sc := ctx.Value(streamCtxKey{}).(streamCtx)
	err := sc.stream.Send(&Downlink{
		T: MsgType_PUSH,
		Push: &ServerPush{
			Evt: msg.Body().(string),
		},
	})
	if err != nil {
		sc.cancel()
		return false, fmt.Errorf("server push %q: stream send: %v", sc.account.ID, err)
	}
	return true, nil
}
