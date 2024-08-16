package tests

import (
	"task-manager/domain"
	"task-manager/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskRepositoryTestSuite struct {
	suite.Suite
	mockRepo *mocks.MockTaskRepository
}

func (suite *TaskRepositoryTestSuite) SetupTest() { // Correct method name
	suite.mockRepo = new(mocks.MockTaskRepository)
}

func (suite *TaskRepositoryTestSuite) TestCreateTask() {
	task := &domain.Task{Title: "new task", Description: "task description"}
	suite.mockRepo.On("CreateTask", task).Return(nil)

	err := suite.mockRepo.CreateTask(task)
	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskRepositoryTestSuite) TestGetAllTasks() {
	tasks := []domain.Task{
		{Title: "Task1", Description: "Description1"},
	}
	suite.mockRepo.On("GetAllTasks").Return(tasks, nil)

	result, err := suite.mockRepo.GetAllTasks()
	suite.NoError(err)
	suite.Equal(tasks, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskRepositoryTestSuite) TestGetTaskByID() {
	taskID := primitive.NewObjectID() // Use consistent ObjectID
	task := &domain.Task{ID: taskID, Title: "Task1", Description: "Description1"}
	suite.mockRepo.On("GetTaskByID", taskID.Hex()).Return(task, nil)

	result, err := suite.mockRepo.GetTaskByID(taskID.Hex())
	suite.NoError(err)
	suite.Equal(task, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskRepositoryTestSuite) TestUpdateTask() {
	task := &domain.Task{ID: primitive.NewObjectID(), Title: "Updated Task", Description: "Updated description"}
	suite.mockRepo.On("UpdateTask", task).Return(nil)

	err := suite.mockRepo.UpdateTask(task)
	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskRepositoryTestSuite) TestDeleteTask() {
	taskID := primitive.NewObjectID().Hex() // Ensure consistent ObjectID usage
	suite.mockRepo.On("DeleteTask", taskID).Return(nil)

	err := suite.mockRepo.DeleteTask(taskID)
	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func TestTaskRepositorySuite(t *testing.T) { // Renamed for clarity
	suite.Run(t, new(TaskRepositoryTestSuite))
}
