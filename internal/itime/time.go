package itime

import "time"

func MakeTimestamp(tm time.Time) int64 {
	if tm.IsZero() {
		return 0
	}
	return tm.UnixNano() / int64(time.Millisecond)
}
