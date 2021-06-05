package rmgr

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

var (
	ErrBadInput = errors.New("bad input parameter")
	ErrExisted  = errors.New("already existed")
	ErrNotFound = errors.New("not found")
	ErrRoomDrop = errors.New("room is dropping")
)

// World is the default room (identity) where a new session starts
const World = ""

// Session is a user's session
type Session struct {
	// user identity
	uid string
	// which room the session belongs
	rid string
	// session birth timestamp
	birth int64
	// last heartbeat timestamp
	heartbeat int64
	// call to cancel the subscription
	cancel context.CancelFunc
}

// Room is a (virtual) place where session gathers
type Room struct {
	sync.RWMutex
	// room identity
	rid string
	// sessions map from user's identity to its session
	sessions map[string]*Session
	// the room is in the dropping phase
	dropping bool
}

// Bucket is a shard of session holder
type Bucket struct {
	sync.RWMutex
	// sessions map from user's identity to its session
	sessions map[string]*Session
	// rooms map from room's identity to the room
	rooms map[string]*Room
}

func NewBucket() *Bucket {
	return &Bucket{
		sessions: make(map[string]*Session),
		rooms:    map[string]*Room{},
	}
}

func newRoom(rid string) *Room {
	return &Room{
		rid:      rid,
		sessions: make(map[string]*Session, 2),
	}
}

// Put a session into the bucket, returns the last session
func (b *Bucket) Put(ses *Session) (*Session, error) {
	errf := func(err error) error {
		return fmt.Errorf("put %v into the bucket: %w", ses, err)
	}
	// check input is valid
	if ses == nil || ses.uid == "" {
		return nil, errf(ErrBadInput)
	}

	var (
		uid  = ses.uid
		room *Room
		ok   bool
		err  error
	)
	b.Lock()
	// check if session already exists
	if s, ok := b.sessions[uid]; ok {
		b.Unlock()
		return s, errf(ErrExisted)
	}

	// now to put a brand new session
	b.sessions[uid] = ses
	// does it also specify a room
	if ses.rid != World {
		// create room if necessary
		if room, ok = b.rooms[ses.rid]; !ok {
			room = newRoom(ses.rid)
			b.rooms[ses.rid] = room
		}
	}
	b.Unlock()

	if room != nil {
		if roomErr := room.put(ses); roomErr != nil {
			err = errf(roomErr)
		}
	}
	return nil, err
}

// JoinRoom join the user's session (uid) to the room (rid), returns the original room (orid)
func (b *Bucket) JoinRoom(uid, rid string) (orid string, err error) {
	errf := func(err error) error {
		return fmt.Errorf("join %q to %q: %w", uid, rid, err)
	}
	// check input is valid
	if uid == "" {
		return "", errf(ErrBadInput)
	}

	var (
		ses  *Session
		room *Room
		ok   bool
	)
	ses = b.FindSession(uid)
	if ses == nil {
		return "", errf(ErrNotFound)
	}
	orid = ses.rid
	// do nothing if room is not changed
	if orid == rid {
		return
	}

	room = b.FindRoom(orid)
	// quit from the original room
	if room != nil && room.del(ses) {
		b.delRoom(orid)
	}
	// if the target room is "World", end
	ses.rid = rid
	if ses.rid == World {
		return
	}
	// create room if necessary
	b.Lock()
	if room, ok = b.rooms[ses.rid]; !ok {
		room = newRoom(ses.rid)
		b.rooms[ses.rid] = room
	}
	b.Unlock()

	if room != nil {
		if roomErr := room.put(ses); roomErr != nil {
			err = errf(roomErr)
		}
	}
	return orid, err
}

func (b *Bucket) FindSession(uid string) *Session {
	b.RLock()
	ses := b.sessions[uid]
	b.RUnlock()
	return ses
}

func (b *Bucket) FindRoom(rid string) *Room {
	b.RLock()
	room := b.rooms[rid]
	b.RUnlock()
	return room
}

func (b *Bucket) delRoom(rid string) {
	var room *Room
	b.Lock()
	if room = b.rooms[rid]; room != nil {
		delete(b.rooms, rid)
	}
	b.Unlock()
	if room != nil {
		room.close()
	}
}

func (b *Bucket) DelSession(uid string) {
	var (
		ok   bool
		ses  *Session
		room *Room
	)
	b.Lock()
	if ses, ok = b.sessions[uid]; ok {
		delete(b.sessions, uid)
		if ses.rid != World {
			room = b.rooms[ses.rid]
		}
	}
	b.Unlock()
	if room != nil && room.del(ses) {
		b.delRoom(ses.rid)
	}
}
