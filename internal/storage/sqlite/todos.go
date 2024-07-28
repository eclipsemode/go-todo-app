package sqlite

import (
	"database/sql"
	"fmt"
	"github.com/eclipsemode/go-todo-app/internal/domain/models"
	uuidv4 "github.com/google/uuid"
)

//go:generate mockery --name TodoService
type TodoRepo interface {
	CreateTodo(title string, description string) (string, error)
	GetAllTodos() ([]models.Todo, error)
	GetTodoById(id string) (models.Todo, error)
}

// CreateTodo creates new element in storage
func (s *Storage) CreateTodo(title string, description string) (string, error) {
	const op = "storage.sqlite.CreateTodo"

	uuid, err := uuidv4.NewRandom()
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := s.db.Prepare("INSERT INTO todos (id, title, description) VALUES (?, ?, ?)")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(uuid.String(), title, description)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return uuid.String(), nil
}

// GetAllTodos returns all todos contained in storage
func (s *Storage) GetAllTodos() ([]models.Todo, error) {
	const op = "storage.sqlite.GetAllTodos"
	var todos []models.Todo

	stmt, err := s.db.Prepare("SELECT * FROM todos")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			fmt.Printf("%s: %s\n", op, err)
			return
		}
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (s *Storage) GetTodoById(id string) (models.Todo, error) {
	const op = "storage.sqlite.GetTodo"

	stmt, err := s.db.Prepare("SELECT * FROM todos WHERE id=?")
	if err != nil {
		return models.Todo{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRow(id)

	defer func() {
		if err := stmt.Close(); err != nil {
			fmt.Printf("%s: %s\n", op, err)
			return
		}
	}()

	var todo models.Todo

	if err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
		return models.Todo{}, fmt.Errorf("%s: %w", op, err)
	}

	return todo, nil
}
