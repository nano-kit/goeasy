// +build linux freebsd darwin

package redir

import (
	"os"
	"path/filepath"
	"syscall"

	"github.com/micro/go-micro/v2/logger"
	ipath "github.com/nano-kit/goeasy/internal/path"
)

func RedirectStdoutStderrToFile(name string, enable bool) {
	if !enable {
		return
	}

	dir := filepath.Join(ipath.HomeDir(), ".microstd")
	os.MkdirAll(dir, 0700)

	if file, err := os.OpenFile(filepath.Join(dir, name+"_out.log"), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666); err == nil {
		stdoutFileHandler = file
		if err := syscall.Dup2(int(file.Fd()), int(os.Stdout.Fd())); err != nil {
			logger.Warn(err)
		}
	} else {
		logger.Warn(err)
	}

	if file, err := os.OpenFile(filepath.Join(dir, name+"_err.log"), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666); err == nil {
		stderrFileHandler = file
		if err := syscall.Dup2(int(file.Fd()), int(os.Stderr.Fd())); err != nil {
			logger.Warn(err)
		}
	} else {
		logger.Warn(err)
	}
}
