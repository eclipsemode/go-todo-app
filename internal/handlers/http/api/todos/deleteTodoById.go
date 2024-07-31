package todos

import (
	responseApi "github.com/eclipsemode/go-todo-app/internal/lib/api/response"
	"github.com/eclipsemode/go-todo-app/internal/lib/logger/sl"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type deleteTodoByIdRes struct {
	responseApi.Response
}

// DeleteTodoById godoc
//
//	@Summary		Delete To-do by id
//	@Description	delete to-do by id
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string		true	"Todos id"
//	@Success		204 "Success"
//	@Failure		400		{object}	responseApi.Response			"Error"
//	@Failure		500		{object}	responseApi.Response			"Error"
//	@Router			/todos/{id} [delete]
func (t *TodoHandler) deleteTodoById(ctx *gin.Context) {
	const op = "handlers.todos.DeleteTodoById"

	id := ctx.Param("id")

	log := t.Log.With(
		slog.String("operation", op),
		slog.String("request_id", requestid.Get(ctx)),
	)

	err := t.UCase.DeleteTodoById(id)
	if err != nil {
		log.Error("cannot delete todo", sl.Err(err))

		ctx.IndentedJSON(http.StatusBadRequest, responseApi.Error("failed to delete todo"))

		return
	}

	log.Info("todo deleted", slog.String("id", id))

	responseDeleteTodoById(ctx)
}

func responseDeleteTodoById(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusNoContent, deleteTodoByIdRes{
		Response: responseApi.Ok(),
	})
}
