package mocks

import (
	"task-manager/domain"

	"github.com/stretchr/testify/mock"
)

type MockUserUseCase struct {
	mock.Mock
}

func (m *MockUserUseCase) RegisterUser(username, password string) (*domain.User, error) {
	args := m.Called(username, password)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserUseCase) LoginUser(username, password string) (string, error) {
	args := m.Called(username, password)
	return args.String(0), args.Error(1)
}

func (m *MockUserUseCase) PromoteUser(userID string) error {
	args := m.Called(userID)
	return args.Error(0)
}
