package tests

import (
	usecases "task-manager/Usecases"
	"task-manager/domain"
	"task-manager/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UseCaseTestSuite struct {
	suite.Suite
	taskUseCase usecases.TaskUseCase
	mockRepo    *mocks.MockTaskRepository
}

func (suite *UseCaseTestSuite) SetUpTest() {
	suite.mockRepo = new(mocks.MockTaskRepository)
	suite.taskUseCase = usecases.NewTaskUseCase(suite.mockRepo)

}

func (suite *UseCaseTestSuite) TestCreateTask() {
	task := domain.Task{
		Title:       "New Task",
		Description: "Task description",
	}

	suite.mockRepo.On("CreateTask", &task).Return(nil)
	err := suite.taskUseCase.CreateTask(&task)
	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())

}

func (suite *UseCaseTestSuite) TestGetTasks() {
	expectedTasks := []domain.Task{
		{Title: "Task1", Description: "Description1"},
		{Title: "Task2", Description: "Description2"},
	}

	suite.mockRepo.On("GetTasks").Return(expectedTasks, nil)
	tasks, err := suite.taskUseCase.GetAllTasks()
	suite.NoError(err)
	suite.Equal(expectedTasks, tasks)
	suite.mockRepo.AssertExpectations(suite.T())
}

func TestUseCaseSuite(t *testing.T) {
	suite.Run(t, new(UseCaseTestSuite))
}
