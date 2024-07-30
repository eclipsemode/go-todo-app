package app

import (
	"errors"
	"fmt"
	"github.com/eclipsemode/go-todo-app/internal/handlers"
	"github.com/eclipsemode/go-todo-app/internal/lib/logger/sl"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func (a *App) StartHttpServer() error {
	err := a.startHttpServer()
	if err != nil {
		return err
	}

	a.gracefulShutdown(a.Logger)

	return err
}

func (a *App) startHttpServer() error {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(requestid.New())

	_, err := handlers.RouterGroup(r, a.Logger, a.Db)
	if err != nil {
		a.Logger.Error("failed to init api v1", sl.Err(err))
		return fmt.Errorf("failed to init api v1: %w", err)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	a.Logger.Info("starting server", slog.Any("config", a.Cfg))

	srv := &http.Server{
		Addr:         a.Cfg.HTTPServer.Addr,
		Handler:      r,
		ReadTimeout:  a.Cfg.HTTPServer.Timeout,
		WriteTimeout: a.Cfg.HTTPServer.Timeout,
		IdleTimeout:  a.Cfg.HTTPServer.Timeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.Logger.Error("error starting server", slog.Any("config", a.Cfg))
		}
	}()

	return nil
}

func (a *App) gracefulShutdown(log *slog.Logger) {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sign := <-quit

	log.Info("app successfully stopped", slog.Any("signal", sign.String()))
}
