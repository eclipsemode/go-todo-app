package app

import (
	h "github.com/eclipsemode/go-todo-app/internal/handlers/http"
	"log/slog"
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
	r := h.NewRouter()
	r.
		WithHandler(a.Logger, a.Db).
		WithSwagger()

	a.Logger.Info("starting server", slog.Any("config", a.Cfg))

	s := h.NewServer(a.Cfg, r.Router)
	s.Start(a.Logger)

	return nil
}

func (a *App) gracefulShutdown(log *slog.Logger) {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sign := <-quit

	log.Info("app successfully stopped", slog.Any("signal", sign.String()))
}
