package config

import (
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Reader struct {
	log *zap.Logger
}

func NewReader(log *zap.Logger) Reader {
	return Reader{log: log}
}

func (r Reader) Read() (*viper.Viper, error) {
	// base config
	v := viper.New()
	v.SetConfigFile("./configs/config.toml")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	// env config
	v.SetConfigFile("./configs/config_dev.toml")
	err = v.MergeInConfig()
	if err != nil {
		return nil, err
	}

	// local config
	v.SetConfigFile("./configs/config.local.toml")
	err = v.MergeInConfig()
	if err != nil {
		return nil, err
	}

	// bind and define env vars
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	_ = v.BindEnv("mysql.dsn")

	return v, nil
}
