package env

import (
	"os"

	"github.com/bamboooo-dev/himo-ingame/internal/interface/mysql"
	"github.com/jinzhu/configor"
	"golang.org/x/xerrors"
)

type Config struct {
	HimoMySQL mysql.Config `yaml:"himoMySQL"`
}

// configPath is tmp file of config
const configPath = "config.yml"

// LoadConfigFromTemplate はテンプレートから設定を読み込む関数
func LoadConfigFromTemplate() (config *Config, err error) {
	config = new(Config)
	err = generateYamlFromTemplate(configPath)
	if err != nil {
		return
	}

	defer func() {
		removeErr := os.Remove(configPath)
		if removeErr != nil {
			err = xerrors.Errorf("%v: %w", removeErr, err)
		}
	}()

	err = configor.Load(config, configPath)
	if err != nil {
		return
	}

	return
}
