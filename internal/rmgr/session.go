package rmgr

import "time"

func (s *Session) Close() {
	if s.cancel != nil {
		s.cancel()
	}
}

func (s *Session) UID() string {
	return s.uid
}

func (s *Session) RID() string {
	return s.rid
}

func (s *Session) Birth() string {
	return makeTime(s.birth).Format(time.StampMilli)
}

func (s *Session) Heartbeat() string {
	return makeTime(s.heartbeat).Format(time.StampMilli)
}
