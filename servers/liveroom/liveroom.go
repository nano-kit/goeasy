package liveroom

import (
	"context"
	"encoding/json"
	math "math"
	"strconv"
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

func roomSequenceKey(room string) string {
	return "room:" + room + ":sequence"
}

// nextSequence 对指定房间的序列号，加1之后返回。每个房间的序列号，从1开始。
func (r *Room) nextSequence(room string) (seq uint64, err error) {
	if room == "" {
		return 0, ierr.BadRequest("empty room identity")
	}
	key := roomSequenceKey(room)
	val, err := r.redisDB.Incr(context.TODO(), key).Result()
	if err != nil {
		return 0, ierr.Storage("INCR %q: %v", key, err)
	}
	return uint64(val), nil
}

// maxSequence 对请求的序列号，返回已分配的最大值。
func (r *Room) maxSequence(room string) (max uint64, err error) {
	if room == "" {
		return 0, ierr.BadRequest("empty room identity")
	}
	key := roomSequenceKey(room)
	val, err := r.redisDB.Get(context.TODO(), key).Result()
	if err != nil {
		return 0, ierr.Storage("GET %q: %v", key, err)
	}
	max, err = strconv.ParseUint(val, 10, 64)
	if err != nil {
		return 0, ierr.Storage("ParseUint GET %q: %v", key, err)
	}
	return max, nil
}

func roomMessageKey(room string) string {
	return "room:" + room + ":messages"
}

// saveRoomMessage 保存到房间消息列表
func (r *Room) saveRoomMessage(msg *RoomMessage) error {
	if msg.Room == "" {
		return ierr.BadRequest("empty room identity")
	}
	key := roomMessageKey(msg.Room)
	data, err := json.Marshal(msg)
	if err != nil {
		return ierr.Internal("marshal room message: %v", err)
	}
	add, err := r.redisDB.ZAdd(context.TODO(), key, &redis.Z{
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
	return err
}

// readRoomMessage 读房间消息
func (r *Room) readRoomMessage(room string, min uint64) ([]*RoomMessage, error) {
	if room == "" {
		return nil, ierr.BadRequest("empty room identity")
	}
	key := roomMessageKey(room)
	data, err := r.redisDB.ZRangeByScore(context.TODO(), key, &redis.ZRangeBy{
		Min: strconv.FormatUint(min, 10),
		Max: "+inf",
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

func roomUpdateKey(room string) string {
	return "room." + room + ".update"
}

// waitForNewRoomMessage 等待有新的房间消息
func (r *Room) waitForNewRoomMessage(room string, timeout time.Duration) error {
	if room == "" {
		return ierr.BadRequest("empty room identity")
	}
	if timeout < 10*time.Second {
		return ierr.BadRequest("short polling timeout")
	}
	sub, err := r.natsConn.SubscribeSync(roomUpdateKey(room))
	if err != nil {
		return ierr.Internal("subscribe: %v", err)
	}
	if err := sub.AutoUnsubscribe(1); err != nil {
		return ierr.Internal("auto unsubscribe: %v", err)
	}
	// Wait for a message
	if _, err = sub.NextMsg(timeout); err == nats.ErrTimeout {
		return ierr.Timeout("no new message for room %q after %vs", room, math.Round(timeout.Seconds()))
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
	seq, err := r.nextSequence(req.Room)
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
	err = r.saveRoomMessage(msg)
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
	max, err := r.maxSequence(req.Room)
	if err != nil {
		return err
	}
	// 有新消息
	if req.LastSeq < max {
		res.Msgs, err = r.readRoomMessage(req.Room, req.LastSeq+1)
		return err
	}
	// 等待
	deadline, ok := ctx.Deadline()
	if !ok {
		return ierr.BadRequest("no deadline")
	}
	if err := r.waitForNewRoomMessage(req.Room, time.Until(deadline.Add(-time.Second))); err != nil {
		return err
	}
	// 重放请求
	res.Msgs, err = r.readRoomMessage(req.Room, req.LastSeq+1)
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
	if err := validateRoom(acc.ID, req.Room); err != nil {
		return err
	}
	seq, err := r.nextSequence(req.Room)
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
	err = r.saveRoomMessage(msg)
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
	seq, err := r.nextSequence(req.Room)
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
	err = r.saveRoomMessage(msg)
	return err
}
