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

func NewTodoHandler(rg *gin.RouterGroup, svc TodoService, log *slog.Logger) error {
	handler := &TodoHandler{
		Service: svc,
		Log:     log,
	}
	rg.POST("/todos", handler.CreateTodoHandler)
	rg.GET("/todos", handler.GetAllTodosHandler)

	return nil
}
