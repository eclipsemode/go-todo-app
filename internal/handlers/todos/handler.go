package todos

import (
	"github.com/eclipsemode/go-todo-app/internal/storage/sqlite"
	"github.com/eclipsemode/go-todo-app/internal/usecase/todos"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type TodoHandler struct {
	uCase *uc_todos.Usecase
	Log   *slog.Logger
}

func NewTodoHandler(rg *gin.RouterGroup, svc sqlite.TodoRepo, log *slog.Logger) error {

	ucaseTodos := uc_todos.New(svc)

	handler := &TodoHandler{
		uCase: ucaseTodos,
		Log:   log,
	}
	rg.POST("/todos", handler.CreateTodoHandler)
	rg.GET("/todos", handler.GetAllTodosHandler)
	rg.GET("/todos/:id", handler.GetTodoById)
	rg.DELETE("/todos/:id", handler.DeleteTodoById)

	return nil
}
