package path

import "os"

func HomeDir() string {
	if dir, err := os.UserHomeDir(); err == nil {
		return dir
	}
	return os.TempDir()
}
