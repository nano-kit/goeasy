package impl

import (
	"context"

	"github.com/nano-kit/goeasy/internal/proto"
)

type User struct {
}

func (u *User) onUserActivity(ctx context.Context, event *proto.UserActivityEvent) error {

	return nil
}
