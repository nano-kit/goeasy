// +build linux

package rlimit

import (
	"golang.org/x/sys/unix"
)

// SetNumFiles sets the linux rlimit for the maximum open files.
func SetNumFiles(maxOpenFiles uint64) error {
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &unix.Rlimit{Max: maxOpenFiles, Cur: maxOpenFiles})
}
