package handlers

import (
	"fmt"
	"github.com/eclipsemode/go-todo-app/internal/handlers/todos"
	"github.com/eclipsemode/go-todo-app/internal/lib/logger/sl"
	"github.com/eclipsemode/go-todo-app/internal/storage/sqlite"
	"github.com/gin-gonic/gin"
	"log/slog"
)

func RouterGroup(r *gin.Engine, log *slog.Logger, storage *sqlite.Storage) (*gin.RouterGroup, error) {
	apiV1 := r.Group("/api/v1")
	{
		err := todos.NewTodoHandler(apiV1, storage, log)
		if err != nil {
			log.Error("failed to init todo handler", sl.Err(err))

			return nil, fmt.Errorf("failed to init todo handler: %w", err)
		}
	}
	return apiV1, nil
}
