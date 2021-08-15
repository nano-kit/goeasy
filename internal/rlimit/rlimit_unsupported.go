// +build !linux

package rlimit

import (
	"errors"
)

// SetNumFiles sets the rlimit for the maximum open files.
func SetNumFiles(maxOpenFiles uint64) error {
	return errors.New("SetRLimit unsupported in this platform")
}
