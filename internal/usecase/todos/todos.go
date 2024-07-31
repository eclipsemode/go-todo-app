package ucTodos

import (
	"github.com/eclipsemode/go-todo-app/internal/domain/models"
	"github.com/eclipsemode/go-todo-app/internal/storage/sqlite"
	"github.com/google/uuid"
)

type Usecase struct {
	Repo sqlite.TodoRepo
}

func New(repo sqlite.TodoRepo) *Usecase {
	return &Usecase{Repo: repo}
}

func (uc *Usecase) CreateTodo(title string, description string) (string, error) {
	return uc.Repo.CreateTodo(title, description)
}

func (uc *Usecase) GetAllTodos() ([]models.Todo, error) {
	return uc.Repo.GetAllTodos()
}

func (uc *Usecase) GetTodoById(id string) (models.Todo, error) {
	return uc.Repo.GetTodoById(id)
}

func (uc *Usecase) DeleteTodoById(id string) error {
	return uc.Repo.DeleteTodoById(id)
}

func (uc *Usecase) UpdateTodoById(id uuid.UUID, title string, description string) (uuid.UUID, error) {
	return uc.Repo.UpdateTodoById(id, title, description)
}
