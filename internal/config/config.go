package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"goTgTemplate/pkg/logging"
	"sync"
)

type Config struct {
	IsDebug     bool    `yaml:"is_debug" env-description:"Is debug?" env-required:"true"`
	BotToken    string  `yaml:"bot_token" env-required:"true"`
	DatabaseDsn string  `yaml:"database_dsn" env-description:"Database DSN" env-required:"true"`
	NatsDsn     *string `yaml:"nats_dsn" env-description:"Nats DSN"`
	RedisDsn    *string `yaml:"redis_dsn" env-description:"Redis DSN"`
}

var DefaultConfig *Config
var once sync.Once

func initConfig() *Config {
	once.Do(func() {
		DefaultConfig = &Config{}
		if err := cleanenv.ReadConfig(".config.yml", DefaultConfig); err != nil {
			logging.DefaultLogger.Panic().Err(err)
			panic(err)
		}
	})
	return DefaultConfig
}

func init() {
	DefaultConfig = initConfig()
}
