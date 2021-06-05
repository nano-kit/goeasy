package rmgr

func (s *Session) Close() {
	if s.cancel != nil {
		s.cancel()
	}
}
