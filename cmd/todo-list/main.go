package main

import (
	"errors"
	"github.com/eclipsemode/go-todo-app/internal/config"
	"github.com/eclipsemode/go-todo-app/internal/http-server/handlers/todos"
	"github.com/eclipsemode/go-todo-app/internal/lib/logger/sl"
	"github.com/eclipsemode/go-todo-app/internal/storage/sqlite"
	"github.com/eclipsemode/logger-pretty"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	logDefault "log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logDefault.Fatal("Error loading .env file")
	}

	cfg := config.MustLoad()

	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(requestid.New())

	log := logger_pretty.NewPrettySlog()

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	apiV1 := r.Group("/api/v1")
	{
		err = todos.NewTodoHandler(apiV1, storage, log)
		if err != nil {
			log.Error("failed to init todo handler", sl.Err(err))

			return
		}
	}

	log.Info("starting server", slog.Any("config", cfg))

	srv := &http.Server{
		Addr:         cfg.HTTPServer.Addr,
		Handler:      r,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.Timeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("error starting server", slog.Any("config", cfg))
		}
	}()

	gracefulShutdown(log)
}

func gracefulShutdown(log *slog.Logger) {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sign := <-quit

	log.Info("app successfully stopped", slog.Any("signal", sign.String()))
}
