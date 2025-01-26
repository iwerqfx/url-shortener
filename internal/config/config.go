package config

import (
	"context"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sethvargo/go-envconfig"
	"sync"
)

type (
	Config struct {
		Log LogConfig
	}

	LogConfig struct {
		Level  string `env:"LOG_LEVEL, required"`
		Format string `env:"LOG_FORMAT, required"`
	}
)

var (
	instance *Config
	once     sync.Once
)

func Get() *Config {
	once.Do(func() {
		var cfg Config
		if err := envconfig.Process(context.Background(), &cfg); err != nil {
			panic("error loading config: " + err.Error())
		}

		instance = &cfg
	})

	return instance
}
