package todos

import (
	"github.com/eclipsemode/go-todo-app/internal/domain/models"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type TodoService interface {
	CreateTodo(title string, description string) (string, error)
	GetAllTodos() ([]models.Todo, error)
}

type TodoHandler struct {
	Service TodoService
	Log     *slog.Logger
}

func NewTodoHandler(r *gin.Engine, svc TodoService, log *slog.Logger) error {
	handler := &TodoHandler{
		Service: svc,
		Log:     log,
	}
	r.POST("/todos", handler.CreateTodoHandler)
	r.GET("/todos", handler.GetAllTodosHandler)

	return nil
}
