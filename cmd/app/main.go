package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/iwerqfx/url-shortener/internal/config"
	"github.com/iwerqfx/url-shortener/internal/handler"
	"github.com/iwerqfx/url-shortener/internal/logger"
	"github.com/iwerqfx/url-shortener/internal/repository/sqlite"
	"github.com/iwerqfx/url-shortener/internal/service"
	slogchi "github.com/samber/slog-chi"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	db, err := sqlite.NewDB(cfg.Database.URL)
	if err != nil {
		panic(err)
	}

	r := sqlite.NewRepository(log, db)
	urlRepository := sqlite.NewURLRepository(r)
	s := service.NewService(log)
	urlService := service.NewURLService(s, urlRepository)
	h := handler.NewHandler(log)
	urlHandler := handler.NewURLHandler(h, urlService, cfg.Server.Address)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(slogchi.New(log))
	router.Use(middleware.Recoverer)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	router.Post("/", urlHandler.Create)
	router.Get("/{alias}", urlHandler.Redirect)

	log.Info("Starting server", slog.String("address", cfg.Server.Address))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr:         cfg.Server.Address,
		Handler:      router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil {
			log.Error("error starting server")
		}
	}()

	log.Info("Server was started", slog.String("address", srv.Addr))

	<-done
	log.Info("Stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
	defer cancel()

	if err = srv.Shutdown(ctx); err != nil {
		log.Error("error shutting down server")
		return
	}

	log.Info("Server was gracefully shutdown")

	if err = db.Close(); err != nil {
		panic("error closing db: " + err.Error())
	}
}
