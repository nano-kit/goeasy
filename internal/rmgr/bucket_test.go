package rmgr

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBucketPut_BadSession(t *testing.T) {
	b := NewBucket()

	badSession := &Session{}
	s, err := b.Put(badSession)
	assert.True(t, errors.Is(err, ErrBadInput))
	assert.True(t, s == nil)
}

func TestBucketPut_SessionInWorld(t *testing.T) {
	b := NewBucket()

	firstSession := &Session{uid: "101"}
	s1, err := b.Put(firstSession)
	assert.NoError(t, err)
	assert.True(t, s1 == nil)

	secondSession := &Session{uid: "101"}
	s2, err := b.Put(secondSession)
	assert.True(t, errors.Is(err, ErrExisted))
	assert.True(t, s2 == firstSession)

	thirdSession := &Session{uid: "102"}
	s3, err := b.Put(thirdSession)
	assert.NoError(t, err)
	assert.True(t, s3 == nil)

	assert.True(t, len(b.sessions) == 2)
	assert.True(t, len(b.rooms) == 0)

	b.DelSession("101")
	assert.True(t, len(b.sessions) == 1)

	b.DelSession("102")
	assert.True(t, len(b.sessions) == 0)
}

func TestBucketPut_SessionInRoom(t *testing.T) {
	b := NewBucket()

	firstSession := &Session{uid: "101", rid: "room1"}
	s1, err := b.Put(firstSession)
	assert.NoError(t, err)
	assert.True(t, s1 == nil)

	secondSession := &Session{uid: "102", rid: "room1"}
	s2, err := b.Put(secondSession)
	assert.NoError(t, err)
	assert.True(t, s2 == nil)

	thirdSession := &Session{uid: "103", rid: "room2"}
	s3, err := b.Put(thirdSession)
	assert.NoError(t, err)
	assert.True(t, s3 == nil)

	assert.True(t, len(b.sessions) == 3)
	assert.True(t, len(b.rooms) == 2)

	b.DelSession("101")
	assert.True(t, len(b.sessions) == 2)
	assert.True(t, len(b.rooms) == 2)

	b.DelSession("103")
	assert.True(t, len(b.sessions) == 1)
	assert.True(t, len(b.rooms) == 1)
	assert.True(t, b.FindRoom("room2") == nil)

	b.DelSession("102")
	assert.True(t, len(b.sessions) == 0)
	assert.True(t, len(b.rooms) == 0)
}

func TestBucketJoinRoom(t *testing.T) {
	b := NewBucket()

	firstSession := &Session{uid: "101"}
	s1, err := b.Put(firstSession)
	assert.NoError(t, err)
	assert.True(t, s1 == nil)

	orid, err := b.JoinRoom("101", "room1")
	assert.NoError(t, err)
	assert.Equal(t, World, orid)

	orid, err = b.JoinRoom("101", "room2")
	assert.NoError(t, err)
	assert.Equal(t, "room1", orid)

	secondSession := &Session{uid: "102", rid: "room1"}
	s2, err := b.Put(secondSession)
	assert.NoError(t, err)
	assert.True(t, s2 == nil)

	assert.True(t, len(b.sessions) == 2)
	assert.True(t, len(b.rooms) == 2)

	orid, err = b.JoinRoom("102", "room2")
	assert.NoError(t, err)
	assert.Equal(t, "room1", orid)
	assert.True(t, len(b.sessions) == 2)
	assert.True(t, len(b.rooms) == 1)
	assert.True(t, len(b.FindRoom("room2").sessions) == 2)

	orid, err = b.JoinRoom("102", "room2")
	assert.NoError(t, err)
	assert.Equal(t, "room2", orid)

	b.JoinRoom("101", World)
	b.JoinRoom("102", World)
	assert.True(t, len(b.sessions) == 2)
	assert.True(t, len(b.rooms) == 0)

	orid, err = b.JoinRoom("", "")
	assert.True(t, errors.Is(err, ErrBadInput))
	assert.Equal(t, "", orid)

	orid, err = b.JoinRoom("nobody", "")
	assert.True(t, errors.Is(err, ErrNotFound))
	assert.Equal(t, "", orid)
}
