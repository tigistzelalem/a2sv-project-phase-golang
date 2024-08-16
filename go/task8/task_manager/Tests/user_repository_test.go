package tests

import (
	"task-manager/domain"
	"task-manager/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	mockRepo *mocks.MockUserRepository
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	suite.mockRepo = new(mocks.MockUserRepository)
}

func (suite *UserRepositoryTestSuite) TestCreateUser() {
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "newuser",
		Password: "password123",
	}

	suite.mockRepo.On("CreateUser", user).Return(nil)

	err := suite.mockRepo.CreateUser(user)

	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserRepositoryTestSuite) TestLoginUser() {
	username := "existinguser"
	user := &domain.User{ID: primitive.NewObjectID(), Username: username, Password: "password123"}

	suite.mockRepo.On("LoginUser", username).Return(user, nil)

	result, err := suite.mockRepo.LoginUser(username)

	suite.NoError(err)
	suite.Equal(user, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserRepositoryTestSuite) TestPromoteUser() {
	userID := primitive.NewObjectID().Hex()

	suite.mockRepo.On("PromoteUser", userID).Return(nil)

	err := suite.mockRepo.PromoteUser(userID)

	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
