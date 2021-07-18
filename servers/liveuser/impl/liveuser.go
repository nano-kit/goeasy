package impl

import (
	"context"

	"github.com/micro/go-micro/v2/logger"
	"github.com/nano-kit/goeasy/internal/json"
	"github.com/nano-kit/goeasy/internal/proto"
)

type User struct {
}

func (u *User) onUserActivity(ctx context.Context, event *proto.UserActivityEvent) error {
	logger.Infof("onUserActivity: %v", json.Stringify(event))
	return nil
}
