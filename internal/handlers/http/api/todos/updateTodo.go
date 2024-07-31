package todos

import (
	"errors"
	responseApi "github.com/eclipsemode/go-todo-app/internal/lib/api/response"
	"github.com/eclipsemode/go-todo-app/internal/lib/logger/sl"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
)

type updateTodoReq struct {
	Title       string `json:"title" validate:"required,max=20"`
	Description string `json:"description,omitempty" validate:"max=100"`
}

type updateTodoRes struct {
	responseApi.Response
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
}

// updateTodoHandler godoc
//
//	@Summary		Update To-do by id
//	@Description	update to-do by id
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			message	body		updateTodoReq	true	"Update to-do req"
//	@Param			id	path		string		true	"Todos id"
//	@Success		200	{object}	updateTodoRes "Success"
//	@Failure		400		{object}	responseApi.Response			"Error"
//	@Failure		500		{object}	responseApi.Response			"Error"
//	@Router			/todos/{id} [put]
func (t *TodoHandler) updateTodoHandler(ctx *gin.Context) {
	const op = "handlers.todos.updateTodoHandler"

	log := t.Log.With(
		slog.String("operation", op),
		slog.String("request_id", requestid.Get(ctx)),
	)

	id := ctx.Param("id")

	parsedId, err := uuid.Parse(id)
	if err != nil {
		log.Error("failed to parse todo id", sl.Err(err))
	}

	var req updateTodoReq

	if err := ctx.ShouldBind(&req); err != nil {
		log.Error("failed to bind update todo request", sl.Err(err))

		ctx.IndentedJSON(http.StatusBadRequest, responseApi.Error("failed to bind update todo request"))

		return
	}

	if err := validator.New().Struct(req); err != nil {
		var validateErr validator.ValidationErrors
		errors.As(err, &validateErr)

		log.Error("failed to validate request body", sl.Err(validateErr))

		ctx.IndentedJSON(http.StatusNotAcceptable, responseApi.ValidationError(validateErr))

		return
	}

	resultId, err := t.UCase.UpdateTodoById(parsedId, req.Title, req.Description)
	if err != nil {
		log.Error("failed to update todo", sl.Err(err))

		ctx.IndentedJSON(http.StatusInternalServerError, responseApi.Error(err.Error()))

		return
	}

	log.Info("todo updated successfully", slog.Any("todo", updateTodoRes{
		Id:          resultId,
		Title:       req.Title,
		Description: req.Description,
	}))

	responseUpdatedTodo(ctx, updateTodoRes{
		Id:          resultId,
		Title:       req.Title,
		Description: req.Description,
	})
}

func responseUpdatedTodo(ctx *gin.Context, r updateTodoRes) {
	ctx.IndentedJSON(http.StatusOK, r)
}
