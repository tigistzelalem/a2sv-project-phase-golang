package mocks

import (
	"task-manager/domain"

	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) CreateTask(task *domain.Task) error {
	args := m.Called(task)
	return args.Error(0)

}

func (m *MockTaskRepository) GetAllTasks() ([]domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]domain.Task), args.Error(1)

}

func (m *MockTaskRepository) GetTaskByID(id string) (*domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Task), args.Error(1)
}

func (m *MockTaskRepository) UpdateTask(task *domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
