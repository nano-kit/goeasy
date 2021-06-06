package comet

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/micro/go-micro/v2/auth"
	errs "github.com/micro/go-micro/v2/errors"
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
	// the first message we recv must be auth
	req, err := stream.Recv()
	if err != nil {
		return errs.BadRequest("incorrect-protocol", "stream recv: %v", err)
	}
	token := req.GetAuth().GetToken()
	if req.T != MsgType_AUTH {
		return errs.BadRequest("incorrect-protocol", "expect message type AUTH but got %v: %v", req.T, token)
	}
	account, ok := iauth.AccountFromToken(token)
	if !ok {
		return errs.BadRequest("unidentified-subscriber", "auth token should be JWT: %v", token)
	}
	logger.Infof("subscriber %q enter", account.ID)

	heartbeat := time.Now()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// create a session with account ID and room ID (if there is any)
	ses := rmgr.NewSession(account.ID, req.GetJoin().GetRid(), cancel)
	oses, err := c.rm.Put(ses)
	defer c.rm.DelSession(account.ID)
	if errors.Is(err, rmgr.ErrExisted) {
		logger.Warnf("about to kick out last session: %v", err)
		oses.Close()

		// wait until last session quit
		for c.rm.FindSession(account.ID) != nil {
			time.Sleep(100 * time.Millisecond)
		}
		// now there is no existing session
		_, err = c.rm.Put(ses)
	}
	if err != nil {
		return errs.InternalServerError("broken-session", "%v", err)
	}

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
			return fmt.Errorf("process %q uplink: stream recv: %v", sc.account.ID, err)
		}
		logger.Debugf("RX %q %v", sc.account.ID, uplink)
		// update heartbeat on any uplink message
		*sc.heartbeat = time.Now()
		// handle uplink commands
		switch uplink.T {
		case MsgType_JOIN:
			if _, err = c.rm.JoinRoom(sc.account.ID, uplink.GetJoin().GetRid()); err != nil {
				return fmt.Errorf("process %q uplink: %v", sc.account.ID, err)
			}
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

	// kick out if idle for a long time
	idle := time.Since(*sc.heartbeat)
	if idle > 3*heartbeatDuration {
		return fmt.Errorf("server tick %q: heartbeat delays", sc.account.ID)
	}

	// request uplink probe by sending a downlink heartbeat
	if idle > heartbeatDuration {
		if err := sc.stream.Send(&Downlink{
			T:  MsgType_HB,
			Hb: &Heartbeat{},
		}); err != nil {
			return fmt.Errorf("server tick %q: send heartbeat: %v", sc.account.ID, err)
		}
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

func (c *Comet) Broadcast(ctx context.Context, req *BroadcastReq, res *BroadcastRes) error {
	// world broadcast
	if req.Rid == rmgr.World {
		c.rm.Enumerate(func(ses *rmgr.Session) {
			p := &PublishReq{
				Uid: ses.UID(),
				Evt: req.Evt,
			}
			c.g.Publish(ctx, &serverPush{p})
		})
		return nil
	}

	// room broadcast
	room := c.rm.FindRoom(req.Rid)
	if room == nil {
		return nil
	}
	room.Enumerate(func(ses *rmgr.Session) {
		p := &PublishReq{
			Uid: ses.UID(),
			Evt: req.Evt,
		}
		c.g.Publish(ctx, &serverPush{p})
	})
	return nil
}

func (c *Comet) DumpSession(ctx context.Context, req *DumpSessionReq, res *DumpSessionRes) error {
	sessions := c.rm.SessionsSnapshot()
	world := make([]*Session, 0, len(sessions))
	for _, ses := range sessions {
		world = append(world, &Session{
			Uid:   ses.UID(),
			Rid:   ses.RID(),
			Birth: ses.Birth(),
		})
	}

	rooms := c.rm.RoomsSnapshot()
	rr := make([]*Room, 0, len(rooms))
	for _, room := range rooms {
		rr = append(rr, dumpRoom(room))
	}

	res.World = world
	res.Rooms = rr
	return nil
}

func dumpRoom(room *rmgr.Room) *Room {
	sessions := room.SessionsSnapshot()
	ss := make([]*Session, 0, len(sessions))
	for _, ses := range sessions {
		ss = append(ss, &Session{
			Uid:   ses.UID(),
			Rid:   ses.RID(),
			Birth: ses.Birth(),
		})
	}
	return &Room{Rid: room.RID(), Room: ss}
}
