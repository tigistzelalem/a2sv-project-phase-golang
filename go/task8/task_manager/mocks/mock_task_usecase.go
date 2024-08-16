package mocks

import (
	"task-manager/domain"

	"github.com/stretchr/testify/mock"
)

type MockTaskUseCase struct {
	mock.Mock
}

func (m *MockTaskUseCase) CreateTask(task *domain.Task) error {
	args := m.Called(task)
	return args.Error(0)

}

func (m *MockTaskUseCase) GetAllTasks() ([]domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]domain.Task), args.Error(1)
}
func (m *MockTaskUseCase) GetTaskByID(id string) (*domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Task), args.Error(1)
}

func (m *MockTaskUseCase) UpdateTask(task *domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskUseCase) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
