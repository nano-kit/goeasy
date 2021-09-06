package impl

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/micro/go-micro/v2/auth"
	"github.com/micro/go-micro/v2/logger"
	"github.com/nano-kit/goeasy/internal/ierr"
	ijson "github.com/nano-kit/goeasy/internal/json"
	"github.com/nano-kit/goeasy/internal/proto"
	"github.com/nano-kit/goeasy/servers/liveuser"
	"github.com/uptrace/bun"
)

type User struct {
	redisDB *redis.Client
	sqlDB   *bun.DB
}

type UserRecord struct {
	UID      string `bun:",pk"`
	Name     string
	Agent    string
	Avatar   string
	UpdateAt time.Time
}

func (u *User) Init(sqlDB *bun.DB, redisDB *redis.Client) {
	u.redisDB = redisDB
	u.sqlDB = sqlDB

	// create table
	models := []interface{}{
		(*UserRecord)(nil),
	}
	for _, model := range models {
		if _, err := u.sqlDB.NewCreateTable().
			IfNotExists().
			Model(model).
			Exec(context.TODO()); err != nil {
			logger.Errorf("can not create table: %v", err)
		}
	}
}

func (u *User) onUserActivity(ctx context.Context, event *proto.UserActivityEvent) error {
	logger.Infof("user act %v", ijson.Stringify(event))
	return nil
}

func (u *User) Set(ctx context.Context, req *liveuser.SetUserInfoReq, res *liveuser.SetUserInfoRes) error {
	acc, ok := auth.AccountFromContext(ctx)
	if !ok {
		return ierr.BadRequest("no account")
	}
	if req.User == nil {
		return ierr.BadRequest("user is nil")
	}
	if req.User.Uid == "" {
		req.User.Uid = acc.ID
	}
	if req.User.Uid != acc.ID {
		return ierr.BadRequest("only your account is allowed")
	}
	_, err := u.sqlDB.NewInsert().
		Model(&UserRecord{
			UID:      req.User.Uid,
			Name:     req.User.Name,
			Agent:    req.User.Agent,
			Avatar:   req.User.Avatar,
			UpdateAt: time.Now(),
		}).
		On("CONFLICT (uid) DO UPDATE").
		Set("name = EXCLUDED.name").
		Set("agent = EXCLUDED.agent").
		Set("avatar = EXCLUDED.avatar").
		Set("update_at = EXCLUDED.update_at").
		Exec(ctx)
	if err != nil {
		return ierr.Storage("Set %q: %v", acc.ID, err)
	}
	return nil
}

func (u *User) Get(ctx context.Context, req *liveuser.GetUserInfoReq, res *liveuser.GetUserInfoRes) error {
	acc, ok := auth.AccountFromContext(ctx)
	if !ok {
		return ierr.BadRequest("no account")
	}
	user := &UserRecord{UID: acc.ID}
	err := u.sqlDB.NewSelect().Model(user).WherePK().Scan(ctx)
	if err == sql.ErrNoRows {
		return ierr.NotFound("no such user")
	}
	if err != nil {
		return ierr.Storage("Get %q: %v", acc.ID, err)
	}
	res.User = &liveuser.UserRecord{
		Uid:      user.UID,
		Name:     user.Name,
		Agent:    user.Agent,
		Avatar:   user.Avatar,
		UpdateAt: user.UpdateAt.UnixNano() / 1e6,
	}
	return nil
}
