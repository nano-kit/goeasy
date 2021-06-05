package rmgr

func (s *Session) Close() {
	if s.cancel != nil {
		s.cancel()
	}
}

func (s *Session) UID() string {
	return s.uid
}
