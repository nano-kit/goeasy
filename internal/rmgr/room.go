package rmgr

import (
	"fmt"
	"sort"
)

func (r *Room) put(ses *Session) error {
	errf := func(err error) error {
		return fmt.Errorf("put %v into the room: %w", ses, err)
	}
	r.Lock()
	// check if session already exists
	if _, ok := r.sessions[ses.uid]; ok {
		r.Unlock()
		return errf(ErrExisted)
	}
	if r.dropping {
		r.Unlock()
		return errf(ErrRoomDrop)
	}
	r.sessions[ses.uid] = ses
	r.Unlock()
	return nil
}

func (r *Room) del(ses *Session) bool {
	r.Lock()
	delete(r.sessions, ses.uid)
	r.dropping = len(r.sessions) == 0
	r.Unlock()
	return r.dropping
}

func (r *Room) close() {
	r.RLock()
	for _, ses := range r.sessions {
		ses.Close()
	}
	r.RUnlock()
}

func (r *Room) Enumerate(accept func(*Session)) {
	r.RLock()
	for _, ses := range r.sessions {
		accept(ses)
	}
	r.RUnlock()
}

func (r *Room) SessionsSnapshot() []*Session {
	r.RLock()
	ss := make([]*Session, 0, len(r.sessions))
	for _, ses := range r.sessions {
		ss = append(ss, ses)
	}
	r.RUnlock()
	// sort the sessions based on birth desc
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].birth > ss[j].birth
	})
	return ss
}

func (r *Room) RID() string {
	return r.rid
}
