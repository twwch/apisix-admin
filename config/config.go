package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/twwch/gin-sdk/twlog"
)

type Config struct {
	HttpListen string            `mapstructure:"http_listen"`
	Log        *twlog.LogConf    `mapstructure:"log"`
	ApisixHost string            `mapstructure:"apisix_host"`
	ApisixKey  string            `mapstructure:"apisix_key"`
	MongoConf  map[string]string `mapstructure:"mongo_conf"`
}

var _globalConfig = new(Config)

func Load(file string) (*Config, error) {
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "failed to read config")
	}
	err := viper.Unmarshal(&_globalConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan config")
	}
	return _globalConfig, nil
}

func Get() *Config {
	return _globalConfig
}
