package liveroom

import (
	"context"
	"encoding/json"
	"math"
	"strconv"
	"sync"
	"text/template"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/micro/go-micro/v2/auth"
	"github.com/micro/go-micro/v2/logger"
	"github.com/nano-kit/goeasy/internal/ierr"
	"github.com/nano-kit/goeasy/internal/proto"
	"github.com/nats-io/nats.go"
)

type Room struct {
	redisDB  *redis.Client
	natsConn *nats.Conn

	mu sync.Mutex // protects inflight
	// prevent one account from initiating more than one long-polling session
	inflight map[string]context.CancelFunc
}

func millisecond(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// validateRoom 验证用户是否在房间内
func validateRoom(user, room string) error {
	if user == "" {
		return ierr.BadRequest("empty user identity")
	}
	if room == "" {
		return ierr.BadRequest("empty room identity")
	}
	// TODO check the user has entered into the room
	return nil
}

// validateText 验证字符合法
func validateText(text string) error {
	if html := template.HTMLEscapeString(text); html != text {
		return ierr.BadRequest("text contains html")
	}
	if js := template.JSEscapeString(text); js != text {
		return ierr.BadRequest("text contains js")
	}
	return nil
}

// roomSequenceKey 保存房间的消息序列号生成器
func roomSequenceKey(room string) string {
	return "room:" + room + ":sequence"
}

// nextSequence 对指定房间的序列号，加1之后返回。每个房间的序列号，从1开始。
func (r *Room) nextSequence(ctx context.Context, room string) (seq uint64, err error) {
	if room == "" {
		return 0, ierr.BadRequest("empty room identity")
	}
	key := roomSequenceKey(room)
	val, err := r.redisDB.Incr(ctx, key).Result()
	if err != nil {
		return 0, ierr.Storage("INCR %q: %v", key, err)
	}
	return uint64(val), nil
}

// maxSequence 对请求的序列号，返回已分配的最大值。
func (r *Room) maxSequence(ctx context.Context, room string) (max uint64, err error) {
	if room == "" {
		return 0, ierr.BadRequest("empty room identity")
	}
	key := roomSequenceKey(room)
	val, err := r.redisDB.Get(ctx, key).Result()
	if err != nil {
		return 0, ierr.Storage("GET %q: %v", key, err)
	}
	max, err = strconv.ParseUint(val, 10, 64)
	if err != nil {
		return 0, ierr.Storage("ParseUint GET %q: %v", key, err)
	}
	return max, nil
}

// roomMessageKey 保存房间的消息列表
func roomMessageKey(room string) string {
	return "room:" + room + ":messages"
}

// saveRoomMessage 保存到房间消息列表
func (r *Room) saveRoomMessage(ctx context.Context, msg *RoomMessage) error {
	if msg.Room == "" {
		return ierr.BadRequest("empty room identity")
	}
	key := roomMessageKey(msg.Room)
	data, err := json.Marshal(msg)
	if err != nil {
		return ierr.Internal("marshal room message: %v", err)
	}
	add, err := r.redisDB.ZAdd(ctx, key, &redis.Z{
		Score:  float64(msg.Seq),
		Member: string(data),
	}).Result()
	if err != nil {
		return ierr.Storage("save room message: %v", err)
	}
	if add == 0 {
		logger.Errorf("duplicate message: some message could be lost!")
		err = ierr.Storage("duplicate message")
	}
	r.notifyNewRoomMessage(msg.Room)
	r.delStaleRoomMessage(ctx, msg.Room)
	return err
}

// readRoomMessage 读房间消息
func (r *Room) readRoomMessage(ctx context.Context, room string, min uint64) ([]*RoomMessage, error) {
	if room == "" {
		return nil, ierr.BadRequest("empty room identity")
	}
	key := roomMessageKey(room)
	data, err := r.redisDB.ZRangeByScore(ctx, key, &redis.ZRangeBy{
		Min:   strconv.FormatUint(min, 10),
		Max:   "+inf",
		Count: 100,
	}).Result()
	if err != nil {
		return nil, ierr.Storage("read room message: %v", err)
	}
	messages := make([]*RoomMessage, len(data))
	for i, str := range data {
		msg := new(RoomMessage)
		if err := json.Unmarshal([]byte(str), msg); err != nil {
			logger.Errorf("unmarshal room message: %v", err)
		}
		messages[i] = msg
	}
	return messages, nil
}

// roomUpdateKey 是房间消息有更新的 nats 主题
func roomUpdateKey(room string) string {
	return "room." + room + ".update"
}

// waitForNewRoomMessage 等待有新的房间消息
func (r *Room) waitForNewRoomMessage(ctx context.Context, room string) error {
	if room == "" {
		return ierr.BadRequest("empty room identity")
	}
	deadline, ok := ctx.Deadline()
	if !ok {
		return ierr.BadRequest("no deadline")
	}
	timeout := time.Until(deadline)
	if timeout < 10*time.Second {
		return ierr.BadRequest("short polling timeout")
	}
	sub, err := r.natsConn.SubscribeSync(roomUpdateKey(room))
	if err != nil {
		return ierr.Internal("subscribe: %v", err)
	}
	defer sub.Unsubscribe()
	if err := sub.AutoUnsubscribe(1); err != nil {
		return ierr.Internal("auto unsubscribe: %v", err)
	}
	// Wait for a message
	if _, err = sub.NextMsgWithContext(ctx); err == context.DeadlineExceeded {
		return ierr.PollTimeout("no new message for room %q after %vs", room, math.Round(timeout.Seconds()))
	} else if err == context.Canceled {
		return ierr.Canceled("canceled by a next attempt")
	} else if err != nil {
		return ierr.Internal("next message: %v", err)
	}
	return nil
}

// notifyNewRoomMessage 通知有新的房间消息
func (r *Room) notifyNewRoomMessage(room string) {
	if room == "" {
		return
	}
	if err := r.natsConn.Publish(roomUpdateKey(room), nil); err != nil {
		logger.Errorf("publish: %v", err)
	}
}

func (r *Room) Init() {
	r.inflight = make(map[string]context.CancelFunc)
}

func (r *Room) Send(ctx context.Context, req *SendReq, res *SendRes) error {
	start := time.Now()
	acc, ok := auth.AccountFromContext(ctx)
	if !ok {
		return ierr.BadRequest("no account")
	}
	if err := validateRoom(acc.ID, req.Room); err != nil {
		return err
	}
	if err := validateText(req.Text); err != nil {
		return err
	}
	seq, err := r.nextSequence(ctx, req.Room)
	if err != nil {
		return err
	}
	msg := &RoomMessage{
		Room:   req.Room,
		Seq:    seq,
		Type:   RoomMessage_PLAIN_TEXT,
		Uid:    acc.ID,
		SendAt: millisecond(start),
		PlainText: &MsgPlainText{
			Text: req.Text,
		},
	}
	err = r.saveRoomMessage(ctx, msg)
	return err
}

func (r *Room) Recv(ctx context.Context, req *RecvReq, res *RecvRes) error {
	acc, ok := auth.AccountFromContext(ctx)
	if !ok {
		return ierr.BadRequest("no account")
	}
	if err := validateRoom(acc.ID, req.Room); err != nil {
		return err
	}
	max, err := r.maxSequence(ctx, req.Room)
	if err != nil {
		return err
	}
	// 处理参数
	if req.LastSeq == 0 && req.OffsetNewest {
		req.LastSeq = max
	}
	// 更新用户心跳
	r.updateRoomUser(ctx, req.Room, acc.ID)
	// 有新消息
	if req.LastSeq < max {
		res.Msgs, err = r.readRoomMessage(ctx, req.Room, req.LastSeq+1)
		return err
	}
	// 等待
	deadline, ok := ctx.Deadline()
	if !ok {
		return ierr.BadRequest("no deadline")
	}
	wait := func() error {
		waitCtx, cancel := context.WithDeadline(ctx, deadline.Add(-time.Second))
		defer cancel()
		r.startWait(acc.ID, cancel) // 记录等待状态
		defer r.endWait(acc.ID)
		return r.waitForNewRoomMessage(waitCtx, req.Room)
	}
	if err := wait(); err != nil {
		return err
	}
	// 重放请求
	res.Msgs, err = r.readRoomMessage(ctx, req.Room, req.LastSeq+1)
	return err
}

func (r *Room) onUserActivity(ctx context.Context, event *proto.UserActivityEvent) error {
	return nil
}

func (r *Room) Enter(ctx context.Context, req *EnterReq, res *EnterRes) error {
	start := time.Now()
	acc, ok := auth.AccountFromContext(ctx)
	if !ok {
		return ierr.BadRequest("no account")
	}
	if err := r.addRoomUser(ctx, req.Room, acc.ID); err != nil {
		return err
	}
	seq, err := r.nextSequence(ctx, req.Room)
	if err != nil {
		return err
	}
	msg := &RoomMessage{
		Room:      req.Room,
		Seq:       seq,
		Type:      RoomMessage_ENTER_ROOM,
		Uid:       acc.ID,
		SendAt:    millisecond(start),
		EnterRoom: &MsgEnterRoom{},
	}
	if err = r.saveRoomMessage(ctx, msg); err != nil {
		return err
	}
	res.Uids, err = r.queryRoomUsers(ctx, req.Room)
	return err
}

func (r *Room) Leave(ctx context.Context, req *LeaveReq, res *LeaveRes) error {
	start := time.Now()
	acc, ok := auth.AccountFromContext(ctx)
	if !ok {
		return ierr.BadRequest("no account")
	}
	if err := validateRoom(acc.ID, req.Room); err != nil {
		return err
	}
	r.delRoomUser(ctx, req.Room, acc.ID)
	seq, err := r.nextSequence(ctx, req.Room)
	if err != nil {
		return err
	}
	msg := &RoomMessage{
		Room:      req.Room,
		Seq:       seq,
		Type:      RoomMessage_LEAVE_ROOM,
		Uid:       acc.ID,
		SendAt:    millisecond(start),
		LeaveRoom: &MsgLeaveRoom{},
	}
	err = r.saveRoomMessage(ctx, msg)
	return err
}

// roomUserKey 记录关注这个房间的用户
func roomUserKey(room string) string {
	return "room:" + room + ":users"
}

// userRoomKey 记录用户在关注哪个房间
func userRoomKey(user string) string {
	return "user:" + user + ":rooms"
}

func (r *Room) addRoomUser(ctx context.Context, room, user string) error {
	if room == "" || user == "" {
		return ierr.BadRequest("empty identity")
	}

	r.updateRoomUser(ctx, room, user)

	return nil
}

func (r *Room) delRoomUser(ctx context.Context, room, user string) {
	if room == "" || user == "" {
		return
	}

	r.redisDB.ZRem(ctx, roomUserKey(room), user)
	r.redisDB.ZRem(ctx, userRoomKey(user), room)
}

const roomUserIdleDuration = 125 * time.Second

const roomUserMaxQueryCount = 100

func (r *Room) updateRoomUser(ctx context.Context, room, user string) {
	if room == "" || user == "" {
		return
	}

	now := time.Now()

	r.redisDB.ZAdd(ctx, roomUserKey(room), &redis.Z{
		Score:  float64(millisecond(now)),
		Member: user,
	})
	r.redisDB.Expire(ctx, roomUserKey(room), roomUserIdleDuration)
	r.redisDB.ZAdd(ctx, userRoomKey(user), &redis.Z{
		Score:  float64(millisecond(now)),
		Member: room,
	})
	r.redisDB.Expire(ctx, userRoomKey(user), roomUserIdleDuration)

	r.delStaleRoomUser(ctx, room, user)
}

func (r *Room) delStaleRoomUser(ctx context.Context, room, user string) {
	maxTS := strconv.FormatInt(millisecond(time.Now().Add(-roomUserIdleDuration)), 10)
	r.redisDB.ZRemRangeByScore(ctx, roomUserKey(room), "-inf", maxTS)
	r.redisDB.ZRemRangeByScore(ctx, userRoomKey(user), "-inf", maxTS)
}

func (r *Room) delStaleRoomMessage(ctx context.Context, room string) {
	r.redisDB.ZRemRangeByRank(ctx, roomMessageKey(room), 0, -1001)
}

func (r *Room) queryRoomUsers(ctx context.Context, room string) (uids []string, err error) {
	return r.redisDB.ZRevRangeByScore(ctx, roomUserKey(room), &redis.ZRangeBy{
		Max:   "+inf",
		Min:   "-inf",
		Count: roomUserMaxQueryCount,
	}).Result()
}

func (r *Room) startWait(uid string, cancel context.CancelFunc) {
	r.mu.Lock()
	previousCancel := r.inflight[uid]
	r.inflight[uid] = cancel
	r.mu.Unlock()

	if previousCancel != nil {
		previousCancel()
	}
}

func (r *Room) endWait(uid string) {
	r.mu.Lock()
	delete(r.inflight, uid)
	r.mu.Unlock()
}
