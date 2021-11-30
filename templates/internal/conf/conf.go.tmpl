package conf

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/toolkits/file"
)

var (
	_defaultConfig = "./configs"
	_config        string
)

func init() {
	if os.Getenv("CONFIGS_DIR") != "" {
		_defaultConfig = os.Getenv("CONFIGS_DIR")
	}

	flag.StringVar(&_config, "config", _defaultConfig, "config dirs")
}

// Load load config file
func Load(name string, cfg interface{}) error {
	configPath := filepath.Join(_config, name)
	if !file.IsExist(configPath) {
		return errors.Errorf("configure file (%s) is not exist", configPath)
	}

	configContent, err := file.ToTrimString(configPath)
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(configContent), cfg)
}
