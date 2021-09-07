package comet

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/auth"
	"github.com/micro/go-micro/v2/client"
	errs "github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/util/pubsub"
	iauth "github.com/nano-kit/goeasy/internal/auth"
	"github.com/nano-kit/goeasy/internal/proto"
	"github.com/nano-kit/goeasy/internal/rmgr"
)

// We must hear subscriber's heartbeat within this duration
const heartbeatDuration = 1 * time.Minute

type Comet struct {
	namespace    string
	g            *pubsub.Group
	rm           *rmgr.Bucket
	userActivity micro.Event
}

func NewComet(namespace string, cli client.Client) *Comet {
	return &Comet{
		namespace:    namespace,
		g:            pubsub.New(),
		rm:           rmgr.NewBucket(),
		userActivity: micro.NewEvent(namespace+".topic.user-activity", cli),
	}
}

type streamCtx struct {
	account        *auth.Account
	stream         Comet_SubscribeStream
	cancel         context.CancelFunc
	uplinkActivity *time.Time
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
	if req.Type != Packet_AUTH {
		return errs.BadRequest("incorrect-protocol", "expect message type AUTH but got %v: %v", req.Type, token)
	}

	// TODO inspect the token by go.micro.auth
	account, err := iauth.AccountFromToken(token)
	if err != nil {
		return errs.BadRequest("unidentified-subscriber", "can not extract account from auth token: %v", err)
	}

	// Ensure the accounts issuer matches the namespace being requested
	if account.Issuer != c.namespace {
		return errs.BadRequest("unidentified-subscriber", "Account was not issued by %v", c.namespace)
	}

	md, _ := metadata.FromContext(ctx)
	logger.Infof("subscriber %q enter with meta data: %v", account.ID, md)

	uplinkActivity := time.Now()
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

	// publish user online
	timeOnline := time.Now().UnixNano()
	c.userActivity.Publish(ctx, &proto.UserActivityEvent{
		Type: proto.UserActivityEvent_ONLINE,
		Uid:  account.ID,
		Time: timeOnline,
	})
	if rid := req.GetJoin().GetRid(); rid != "" {
		c.userActivity.Publish(ctx, &proto.UserActivityEvent{
			Type: proto.UserActivityEvent_ENTER_ROOM,
			Uid:  account.ID,
			Rid:  rid,
			Time: timeOnline,
		})
	}

	// the core processing
	ctx = context.WithValue(ctx, streamCtxKey{}, streamCtx{account, stream, cancel, &uplinkActivity})
	c.g.Go(ctx, c.recv)
	c.g.Subscribe(ctx, account.ID, c.send, pubsub.WithTicker(heartbeatDuration, c.tick))

	// publish user offline
	timeOffline := time.Now().UnixNano()
	if rid := ses.RID(); rid != "" {
		c.userActivity.Publish(ctx, &proto.UserActivityEvent{
			Type: proto.UserActivityEvent_LEAVE_ROOM,
			Uid:  account.ID,
			Rid:  rid,
			Time: timeOffline,
		})
	}
	c.userActivity.Publish(ctx, &proto.UserActivityEvent{
		Type: proto.UserActivityEvent_OFFLINE,
		Uid:  account.ID,
		Time: timeOffline,
	})

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
		// update uplink activity on any uplink message
		*sc.uplinkActivity = time.Now()
		// publish user heard
		c.userActivity.Publish(ctx, &proto.UserActivityEvent{
			Type: proto.UserActivityEvent_HEARD,
			Uid:  sc.account.ID,
			Time: sc.uplinkActivity.UnixNano(),
		})
		// handle uplink commands
		switch uplink.Type {
		case Packet_JOIN:
			if orid, err := c.rm.JoinRoom(sc.account.ID, uplink.GetJoin().GetRid()); err != nil {
				return fmt.Errorf("process %q uplink: %v", sc.account.ID, err)
			} else {
				c.publishJoinRoom(ctx, sc.account.ID, uplink.GetJoin().GetRid(), orid)
			}
		}
	}
}

func (c *Comet) publishJoinRoom(ctx context.Context, uid, rid, orid string) {
	now := time.Now().UnixNano()

	// do nothing if room is not changed
	if orid == rid {
		return
	}

	// leave the original room
	if orid != "" {
		c.userActivity.Publish(ctx, &proto.UserActivityEvent{
			Type: proto.UserActivityEvent_LEAVE_ROOM,
			Uid:  uid, Rid: orid, Time: now,
		})
	}

	// enter the target room
	if rid != "" {
		c.userActivity.Publish(ctx, &proto.UserActivityEvent{
			Type: proto.UserActivityEvent_ENTER_ROOM,
			Uid:  uid, Rid: rid, Time: now,
		})
	}
}

func (c *Comet) tick(ctx context.Context) (err error) {
	sc := ctx.Value(streamCtxKey{}).(streamCtx)
	defer func() {
		if err != nil {
			sc.cancel()
		}
	}()

	// kick out if idle for a long time
	idle := time.Since(*sc.uplinkActivity)
	if idle > 3*heartbeatDuration {
		return fmt.Errorf("server tick %q: hasn't been heard of for a long time", sc.account.ID)
	}

	// request uplink probe by sending a downlink heartbeat
	if idle >= heartbeatDuration-time.Second {
		if err := sc.stream.Send(&Downlink{
			Type: Packet_HB,
			Hb:   &Heartbeat{},
		}); err != nil {
			return fmt.Errorf("server tick %q: send heartbeat: %v", sc.account.ID, err)
		}
	}
	return nil
}

func (c *Comet) send(ctx context.Context, msg pubsub.Message) (bool, error) {
	sc := ctx.Value(streamCtxKey{}).(streamCtx)
	err := sc.stream.Send(&Downlink{
		Type: Packet_PUSH,
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

// onEvent will be executed when a event is received.
// Event is published to a topic which any comet instance is subscribing. As
// every user is landing upon a comet, the event will finally reach that user.
// Event that designates to room or world will also reach every user within.
func (c *Comet) onEvent(ctx context.Context, event *proto.Event) error {
	if event.Uid != "" {
		return c.Publish(ctx, &PublishReq{
			Uid: event.Uid,
			Evt: event.Evt,
		}, new(PublishRes))
	}

	return c.Broadcast(ctx, &BroadcastReq{
		Rid: event.Rid,
		Evt: event.Evt,
	}, new(BroadcastRes))
}
