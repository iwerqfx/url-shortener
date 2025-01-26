package main

import (
	"fmt"
	"github.com/iwerqfx/url-shortener/internal/config"
	"github.com/iwerqfx/url-shortener/internal/logger"
)

func main() {
	cfg := config.Get()

	log := logger.MustNew(logger.Config{
		Level:  cfg.Log.Level,
		Format: cfg.Log.Format,
	})

	log.Info(
		fmt.Sprintf("Starting [%s]", cfg.App.Name),
	)

}
