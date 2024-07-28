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

type GetTodoRes struct {
	responseApi.Response
	Todo models.Todo `json:"todo"`
}

// GetTodoById godoc
//
//	@Summary		Get To-do by id
//	@Description	get to-do by id
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string		true	"Todos id"
//	@Success		200	{object}	GetTodoRes "Success"
//	@Failure		400		{object}	responseApi.Response			"Error"
//	@Failure		500		{object}	responseApi.Response			"Error"
//	@Router			/todos/{id} [get]
func (t *TodoHandler) GetTodoById(ctx *gin.Context) {
	const op = "handlers.todos.getTodoById"

	id := ctx.Param("id")

	log := t.Log.With(
		slog.String("operation", op),
		slog.String("request_id", requestid.Get(ctx)),
	)

	todo, err := t.uCase.GetTodoById(id)
	if err != nil {
		log.Error("error getting todo", sl.Err(err))

		ctx.IndentedJSON(http.StatusBadRequest, responseApi.Error("failed to get todo"))

		return
	}

	log.Info("todo successfully received", slog.Any("todo", todo))

	responseGetTodoById(ctx, todo)

}

func responseGetTodoById(ctx *gin.Context, todo models.Todo) {
	ctx.IndentedJSON(http.StatusOK, GetTodoRes{
		Response: responseApi.Ok(),
		Todo:     todo,
	})
}
