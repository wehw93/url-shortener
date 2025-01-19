package main

import (
	"log/slog"
	"net/http"
	"os"
	"url-shortener/internal/config"
	"url-shortener/internal/http-server/handlers/url/deleteURL"
	"url-shortener/internal/http-server/handlers/url/redirect"
	"url-shortener/internal/http-server/handlers/url/save"
	mvLogger "url-shortener/internal/http-server/middleware/logger"
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

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(mvLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Route("/url",func (r chi.Router){
		r.Use(middleware.BasicAuth("url-shortener", map[string]string{
			cfg.HTTPServer.User:cfg.HTTPServer.Password,
		}))
		r.Post("/", save.New(log, storage))
		r.Delete("/{alias}", deleteURL.New(log, storage))
	})

	
	router.Get("/{alias}", redirect.New(log, storage))
	
	
	log.Info("starting server", slog.String("adress", cfg.Addres))

	srv := &http.Server{
		Addr:         cfg.Addres,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.Idle_Timeout,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}
	log.Info("server stopped")

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
