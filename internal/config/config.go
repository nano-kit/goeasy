package config

import (
	microconfig "github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
)

// LoadInitialConfigFromFile loads from JSON or YAML file and unmarshal
// the file's content into the value pointed to by v. The initial config data
// includes but not limited to serving address, namespace and database
// address, which are parameters that should be known before server starts
// and can not be changed until server stops.
func LoadInitialConfigFromFile(filePath string, v interface{}) error {
	conf, err := microconfig.NewConfig(microconfig.WithSource(
		file.NewSource(file.WithPath(filePath)),
	))
	if err != nil {
		return err
	}

	return conf.Scan(v)
}
