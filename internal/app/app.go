package app

import (
	"errors"
	"github.com/eclipsemode/go-todo-app/internal/config"
	"github.com/eclipsemode/go-todo-app/internal/storage/sqlite"
	"github.com/eclipsemode/logger-pretty"
	"log/slog"
)

type App struct {
	Cfg    *config.Config
	Db     *sqlite.Storage
	Logger *slog.Logger
}

var a *App

func NewApp() (*App, error) {
	conf := config.MustLoad()

	log := logger_pretty.NewPrettySlog()

	storage, err := sqlite.New(conf.StoragePath)
	if err != nil {
		return nil, errors.New("failed to initialize storage")
	}

	app := &App{
		Cfg:    conf,
		Db:     storage,
		Logger: log,
	}

	return app, nil
}

func SetGlobalApp(app *App) {
	a = app
}

func GetGlobalApp() (*App, error) {
	if a == nil {
		return nil, errors.New("global app is not initialized")
	}
	return a, nil
}
