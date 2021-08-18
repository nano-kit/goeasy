package liveroom

import (
	context "context"

	"github.com/go-redis/redis/v8"
	"github.com/nano-kit/goeasy/internal/proto"
	"github.com/nats-io/nats.go"
)

type Room struct {
	redisDB  *redis.Client
	natsConn *nats.Conn
}

func (r *Room) Send(ctx context.Context, req *SendReq, res *SendRes) error {
	return nil
}

func (r *Room) Recv(ctx context.Context, req *RecvReq, res *RecvRes) error {
	return nil
}

func (r *Room) onUserActivity(ctx context.Context, event *proto.UserActivityEvent) error {
	return nil
}

func (r *Room) Enter(ctx context.Context, req *EnterReq, res *EnterRes) error {
	return nil
}

func (r *Room) Leave(ctx context.Context, req *LeaveReq, res *LeaveRes) error {
	return nil
}
