package main

import (
	"database/sql"
	"fmt"
	"github.com/iwerqfx/url-shortener/internal/config"
	"github.com/iwerqfx/url-shortener/internal/logger"
	"github.com/iwerqfx/url-shortener/internal/repository/sqlite"
	"github.com/iwerqfx/url-shortener/internal/service"
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

	db, err := sqlite.NewDB(cfg.DB.URL)
	if err != nil {
		panic(err)
	}

	defer func(db *sql.DB) {
		if err = db.Close(); err != nil {
			panic("error closing db: " + err.Error())
		}
	}(db)

	r := sqlite.NewRepository(log, db)
	urlRepository := sqlite.NewURLRepository(r)

	s := service.NewService(log)
	urlService := service.NewURLService(s, urlRepository)

}
