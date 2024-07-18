package main

import (
	"github.com/eclipsemode/logger-pretty"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	logDefault "log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"todo-list/internal/config"
	"todo-list/internal/lib/logger/sl"
	"todo-list/internal/storage/sqlite"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logDefault.Fatal("Error loading .env file")
	}

	cfg := config.MustLoad()

	r := gin.Default()

	log := logger_pretty.NewPrettySlog()

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage

	log.Info("starting server", slog.Any("config", cfg))

	srv := &http.Server{
		Addr:         cfg.HTTPServer.Addr,
		Handler:      r,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.Timeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("error starting server", slog.Any("config", cfg))
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sign := <-quit

	log.Info("app successfully stopped", slog.Any("signal", sign.String()))
}
