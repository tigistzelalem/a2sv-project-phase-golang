package mocks

import (
	"task-manager/domain"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)

}

func (m *MockUserRepository) LoginUser(username string) (*domain.User, error) {
	args := m.Called(username)
	return args.Get(0).(*domain.User), args.Error(1)

}

func (m *MockUserRepository) PromoteUser(userID string) error {
	args := m.Called(userID)
	return args.Error(0)
}
