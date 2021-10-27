// Package ierr provides helpers for common errors when implementing service.
package ierr

import (
	"fmt"

	"github.com/micro/go-micro/v2/errors"
)

// Internal should be used with care because it will trigger a client retry
func Internal(format string, a ...interface{}) error {
	return errors.InternalServerError("error-internal", format, a...)
}

// BadRequest is used with invalid request arguments
func BadRequest(format string, a ...interface{}) error {
	return errors.BadRequest("error-bad-request", format, a...)
}

// Storage is used with any database reported errors
func Storage(format string, a ...interface{}) error {
	return errors.New("error-storage", fmt.Sprintf(format, a...), 550)
}

// PollTimeout is used when timeout event happens during long-polling.
// The difference from 408 is that the latter will trigger a client retry.
func PollTimeout(format string, a ...interface{}) error {
	return errors.New("error-poll-timeout", fmt.Sprintf(format, a...), 551)
}

// NotFound is used when the requested resource does not exist
func NotFound(format string, a ...interface{}) error {
	return errors.NotFound("error-not-found", format, a...)
}

// Canceled is used when the processing is canceled explicitly
func Canceled(format string, a ...interface{}) error {
	return errors.New("error-canceled", fmt.Sprintf(format, a...), 552)
}
