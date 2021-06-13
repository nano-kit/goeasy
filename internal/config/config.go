package config

import (
	"os"
	"path/filepath"
	"strings"

	microconfig "github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
)

// LoadInitialConfigFromFile loads from JSON or YAML file and unmarshal
// the file's content into the value pointed to by v. The initial config data
// includes but not limited to serving address, namespace and database
// address, which are parameters that should be known before server starts
// and can not be changed until server stops.
func LoadInitialConfigFromFile(filePath string, v interface{}) error {
	if !startsWithSlash(filePath) {
		exeDir, err := executableDir()
		if err != nil {
			return err
		}
		filePath = filepath.Join(exeDir, filePath)
	}

	conf, err := microconfig.NewConfig(microconfig.WithSource(
		file.NewSource(file.WithPath(filePath)),
	))
	if err != nil {
		return err
	}

	return conf.Scan(v)
}

func startsWithSlash(filePath string) bool {
	if strings.HasPrefix(filePath, "/") {
		return true
	}
	if filepath.VolumeName(filePath) != "" {
		return true
	}
	return false
}

func executableDir() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return exe, err
	}
	exeDir := filepath.Dir(exe)
	return exeDir, nil
}
