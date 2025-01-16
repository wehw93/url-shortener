package main

import (
	"log/slog"
	"os"
	"url-shortener/internal/config"
	"url-shortener/internal/lib/logger/sl"
	"url-shortener/internal/storage/sqllite"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

const (
	envLocal = "local"
	envDev   = "develop"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	storage, err := sqllite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failrd to init storage", sl.Err(err))
		os.Exit(1)
	}


	router:=chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	//TODO: run server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}

