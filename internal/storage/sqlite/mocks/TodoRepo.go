// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	models "github.com/eclipsemode/go-todo-app/internal/domain/models"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// TodoRepo is an autogenerated mock type for the TodoRepo type
type TodoRepo struct {
	mock.Mock
}

// CreateTodo provides a mock function with given fields: title, description
func (_m *TodoRepo) CreateTodo(title string, description string) (string, error) {
	ret := _m.Called(title, description)

	if len(ret) == 0 {
		panic("no return value specified for CreateTodo")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(title, description)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(title, description)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(title, description)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTodoById provides a mock function with given fields: id
func (_m *TodoRepo) DeleteTodoById(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTodoById")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllTodos provides a mock function with given fields:
func (_m *TodoRepo) GetAllTodos() ([]models.Todo, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllTodos")
	}

	var r0 []models.Todo
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Todo, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Todo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Todo)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTodoById provides a mock function with given fields: id
func (_m *TodoRepo) GetTodoById(id string) (models.Todo, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetTodoById")
	}

	var r0 models.Todo
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.Todo, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) models.Todo); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Todo)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTodoById provides a mock function with given fields: id, title, description
func (_m *TodoRepo) UpdateTodoById(id uuid.UUID, title string, description string) (uuid.UUID, error) {
	ret := _m.Called(id, title, description)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTodoById")
	}

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, string, string) (uuid.UUID, error)); ok {
		return rf(id, title, description)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID, string, string) uuid.UUID); ok {
		r0 = rf(id, title, description)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID, string, string) error); ok {
		r1 = rf(id, title, description)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTodoRepo creates a new instance of TodoRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTodoRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *TodoRepo {
	mock := &TodoRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
