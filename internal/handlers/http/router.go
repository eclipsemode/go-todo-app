package http

import (
	"github.com/eclipsemode/go-todo-app/internal/handlers"
	"github.com/eclipsemode/go-todo-app/internal/lib/logger/sl"
	"github.com/eclipsemode/go-todo-app/internal/storage/sqlite"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log/slog"
)

type Router struct {
	Router *gin.Engine
}

func NewRouter() *Router {
	r := gin.Default()

	r.Use(requestid.New())

	return &Router{Router: r}
}

func (r *Router) WithSwagger() *Router {
	r.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

func (r *Router) WithHandler(log *slog.Logger, storage *sqlite.Storage) *Router {
	_, err := handlers.RouterGroup(r.Router, log, storage)
	if err != nil {
		log.Error("failed to init api v1", sl.Err(err))
	}

	return r
}
