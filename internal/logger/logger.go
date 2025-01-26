package logger

import (
	"errors"
	"github.com/lmittmann/tint"
	"log/slog"
	"os"
	"time"
)

type Config struct {
	Level  string
	Format string
}

func New(cfg Config) (*slog.Logger, error) {
	var handler slog.Handler
	var level slog.Level

	switch cfg.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		return nil, errors.New("invalid log level")
	}

	switch cfg.Format {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	default:
		handler = tint.NewHandler(os.Stdout, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.DateTime,
		})
	}

	return slog.New(handler), nil
}

func MustNew(cfg Config) *slog.Logger {
	logger, err := New(cfg)
	if err != nil {
		panic("error initializing logger: " + err.Error())
	}

	return logger
}
