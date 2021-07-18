package liveroom

import (
	context "context"

	"github.com/nano-kit/goeasy/internal/proto"
)

type Room struct{}

func (r *Room) Send(ctx context.Context, req *SendReq, res *SendRes) error {
	return nil
}

func (r *Room) Recv(ctx context.Context, req *RecvReq, res *RecvRes) error {
	return nil
}

func (r *Room) onUserActivity(ctx context.Context, event *proto.UserActivityEvent) error {
	return nil
}
