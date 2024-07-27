package todos

import (
	"errors"
	responseApi "github.com/eclipsemode/go-todo-app/internal/lib/api/response"
	"github.com/eclipsemode/go-todo-app/internal/lib/logger/sl"
	"github.com/eclipsemode/go-todo-app/internal/storage"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
)

type CreateTodoReq struct {
	Title       string `json:"title" validate:"required,max=20"`
	Description string `json:"description,omitempty" validate:"max=100"`
}

type CreateTodoRes struct {
	responseApi.Response
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

// CreateTodoHandler godoc
//
//	@Summary		Create to-do post method
//	@Description	create to-do
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			message	body		CreateTodoReq	true	"Account Info"
//	@Success		200	{object}	CreateTodoRes "Success"
//	@Failure		400		{object}	responseApi.Response			"Error"
//	@Failure		500		{object}	responseApi.Response			"Error"
//	@Router			/todos [post]
func (t *TodoHandler) CreateTodoHandler(ctx *gin.Context) {
	const op = "handlers.todos.create.New"

	log := t.Log.With(
		slog.String("operation", op),
		slog.String("request_id", requestid.Get(ctx)),
	)

	var req CreateTodoReq

	if err := ctx.ShouldBind(&req); err != nil {
		log.Error("failed to bind the body", sl.Err(err))

		ctx.IndentedJSON(http.StatusBadRequest, responseApi.Error("failed to bind body"))

		return
	}

	log.Info("request body successfully parsed", slog.Any("request", req))

	if err := validator.New().Struct(req); err != nil {
		var validateErr validator.ValidationErrors
		errors.As(err, &validateErr)

		log.Error("failed to validate request body", sl.Err(validateErr))

		ctx.IndentedJSON(http.StatusNotAcceptable, responseApi.ValidationError(validateErr))

		return
	}

	_, err := t.Service.CreateTodo(req.Title, req.Description)
	if errors.Is(err, storage.ErrAlreadyExists) {
		log.Error("todo already exists", sl.Err(err))

		ctx.IndentedJSON(http.StatusInternalServerError, responseApi.Error("todo already exists"))

		return
	}
	if err != nil {
		log.Error("failed to create todo", sl.Err(err))

		ctx.IndentedJSON(http.StatusInternalServerError, responseApi.Error("failed to create todo"))

		return
	}

	log.Info("todo successfully created", slog.Any("todo", req))

	responseCreated(ctx, req.Title, req.Description)
}

func responseCreated(ctx *gin.Context, title string, description string) {
	ctx.IndentedJSON(http.StatusCreated, CreateTodoRes{
		Response:    responseApi.Ok(),
		Title:       title,
		Description: description,
	})
}
