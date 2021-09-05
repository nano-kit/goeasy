package impl

import (
	"context"
	"encoding/json"
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
	UID      string
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

func userRecordKey(uid string) string {
	return "user:" + uid + ":record"
}

func (u *User) AddUser(ctx context.Context, req *liveuser.AddUserReq, res *liveuser.AddUserRes) error {
	acc, ok := auth.AccountFromContext(ctx)
	if !ok {
		return ierr.BadRequest("no account")
	}
	if req.User == nil {
		return ierr.BadRequest("user is nil")
	}
	if req.User.Uid != acc.ID {
		return ierr.BadRequest("only your account is allowed")
	}
	var record map[string]interface{}
	data, _ := json.Marshal(req.User)
	json.Unmarshal(data, &record)
	err := u.redisDB.HSet(ctx, userRecordKey(acc.ID), record).Err()
	if err != nil {
		return ierr.Storage("HSET %q: %v", acc.ID, err)
	}
	u.redisDB.Expire(ctx, userRecordKey(acc.ID), 72*time.Hour)
	return nil
}

func (u *User) QueryUser(ctx context.Context, req *liveuser.QueryUserReq, res *liveuser.QueryUserRes) error {
	if len(req.Uids) == 0 {
		return ierr.BadRequest("empty uids")
	}
	cmds := make([]*redis.StringStringMapCmd, len(req.Uids))
	res.Users = make([]*liveuser.UserRecord, len(req.Uids))
	pipe := u.redisDB.Pipeline()
	for i, uid := range req.Uids {
		cmds[i] = pipe.HGetAll(ctx, userRecordKey(uid))
	}
	if _, err := pipe.Exec(ctx); err != nil {
		return ierr.Storage("HGETALL: %v", err)
	}
	for i := range cmds {
		ssm, _ := cmds[i].Result()
		data, _ := json.Marshal(ssm)
		user := new(liveuser.UserRecord)
		json.Unmarshal(data, user)
		res.Users[i] = user
	}
	return nil
}
