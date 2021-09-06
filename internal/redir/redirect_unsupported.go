// +build !linux,!freebsd,!darwin

package redir

import (
	"github.com/micro/go-micro/v2/logger"
)

func RedirectStdoutStderrToFile(name string, enable bool) {
	logger.Info("Can't redirect stdout and stderr to file")
}
