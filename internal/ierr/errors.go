// Package ierr provides helpers for common errors when implementing service.
package ierr

import (
	"fmt"

	"github.com/micro/go-micro/v2/errors"
)

// Internal should be used with care because it will trigger a client retry
func Internal(format string, a ...interface{}) error {
	return errors.InternalServerError("E0000", format, a...)
}

// BadRequest is used with invalid request arguments
func BadRequest(format string, a ...interface{}) error {
	return errors.BadRequest("E0001", format, a...)
}

// Storage is used with any database reported errors
func Storage(format string, a ...interface{}) error {
	return errors.New("E0100", fmt.Sprintf(format, a...), 550)
}

// Timeout is used with when timeout event happens
func Timeout(format string, a ...interface{}) error {
	return errors.New("E0101", fmt.Sprintf(format, a...), 551)
}
