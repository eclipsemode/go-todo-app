package todos

import (
	"github.com/eclipsemode/go-todo-app/internal/domain/models"
	responseApi "github.com/eclipsemode/go-todo-app/internal/lib/api/response"
	"github.com/eclipsemode/go-todo-app/internal/lib/logger/sl"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type GetAllTodosRes struct {
	responseApi.Response
	Todos []models.Todo `json:"todos"`
}

// GetAllTodosHandler godoc
//
//	@Summary		Show a to-do
//	@Description	get all todos
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	GetAllTodosRes "Success"
//	@Failure		400		{object}	responseApi.Response			"Error"
//	@Failure		500		{object}	responseApi.Response			"Error"
//	@Router			/todos [get]
func (t *TodoHandler) GetAllTodosHandler(ctx *gin.Context) {
	const op = "handlers.todos.getAllTodosHandler"

	log := t.Log.With(
		slog.String("operation", op),
		slog.String("request_id", requestid.Get(ctx)),
	)

	todos, err := t.uCase.GetAllTodos()
	if err != nil {
		log.Error("error getting todos", sl.Err(err))

		ctx.IndentedJSON(http.StatusBadRequest, responseApi.Error("failed to get all todos"))

		return
	}

	log.Info("success getting all todos", slog.Any("todos", todos))

	responseGetAllTodos(ctx, todos)
}

func responseGetAllTodos(ctx *gin.Context, todos []models.Todo) {
	ctx.IndentedJSON(http.StatusOK, GetAllTodosRes{
		Response: responseApi.Ok(),
		Todos:    todos,
	})
}
