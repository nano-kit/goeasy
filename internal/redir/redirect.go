package redir

import "os"

var (
	stdoutFileHandler *os.File //lint:ignore U1000 to avoid gc
	stderrFileHandler *os.File //lint:ignore U1000 to avoid gc
)
