package todos

import (
	"github.com/eclipsemode/go-todo-app/internal/storage/sqlite"
	"github.com/eclipsemode/go-todo-app/internal/usecase/todos"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type TodoHandler struct {
	UCase *ucTodos.Usecase
	Log   *slog.Logger
}

func NewTodoHandler(rg *gin.RouterGroup, svc sqlite.TodoRepo, log *slog.Logger) error {

	uCaseTodos := ucTodos.New(svc)

	handler := &TodoHandler{
		UCase: uCaseTodos,
		Log:   log,
	}
	rg.POST("/todos", handler.createTodoHandler)
	rg.GET("/todos", handler.getAllTodosHandler)
	rg.GET("/todos/:id", handler.getTodoById)
	rg.DELETE("/todos/:id", handler.deleteTodoById)
	rg.PUT("/todos/:id", handler.updateTodoHandler)

	return nil
}
