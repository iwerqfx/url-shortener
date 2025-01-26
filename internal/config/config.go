package config

import (
	"context"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sethvargo/go-envconfig"
	"sync"
	"time"
)

type (
	Config struct {
		App      AppConfig
		Log      LogConfig
		Database DatabaseConfig
		Server   ServerConfig
	}

	AppConfig struct {
		Name string `env:"APP_NAME, required"`
	}

	LogConfig struct {
		Level  string `env:"LOG_LEVEL, required"`
		Format string `env:"LOG_FORMAT, required"`
	}

	DatabaseConfig struct {
		URL string `env:"DATABASE_URL, required"`
	}

	ServerConfig struct {
		Address         string        `env:"SERVER_ADDRESS, required"`
		ReadTimeout     time.Duration `env:"SERVER_READ_TIMEOUT, required"`
		WriteTimeout    time.Duration `env:"SERVER_WRITE_TIMEOUT, required"`
		IdleTimeout     time.Duration `env:"SERVER_IDLE_TIMEOUT, required"`
		ShutdownTimeout time.Duration `env:"SERVER_SHUTDOWN_TIMEOUT, required"`
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
