package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"task-manager/delivery/controllers"
	"task-manager/domain"
	"task-manager/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ControllerTestSuite struct {
	suite.Suite
	controller      *controllers.Controller
	mockUserUseCase *mocks.MockUserUseCase
	mockTaskUseCase *mocks.MockTaskUseCase
	router          *gin.Engine
}

func (suite *ControllerTestSuite) SetupTest() {
	suite.mockUserUseCase = new(mocks.MockUserUseCase)
	suite.mockTaskUseCase = new(mocks.MockTaskUseCase)
	suite.controller = controllers.NewController(suite.mockUserUseCase, suite.mockTaskUseCase)

	suite.router = gin.Default()

	// Register routes for testing
	suite.router.POST("/register", suite.controller.RegisterUser)
	suite.router.POST("/login", suite.controller.LoginUser)
	suite.router.PUT("/promote/:id", suite.controller.PromoteUser)
	suite.router.POST("/tasks", suite.controller.CreateTask)
	suite.router.GET("/tasks", suite.controller.GetAllTasks)
	suite.router.GET("/tasks/:id", suite.controller.GetTaskByID)
	suite.router.PUT("/tasks/:id", suite.controller.UpdateTask)
	suite.router.DELETE("/tasks/:id", suite.controller.DeleteTask)
}

func (suite *ControllerTestSuite) TestCreateTaskHandler() {
	task := domain.Task{
		Title:       "New Task",
		Description: "Task description",
	}

	suite.mockTaskUseCase.On("CreateTask", &task).Return(nil)

	reqBody := bytes.NewBuffer([]byte(`{"title": "New Task", "description": "Task description"}`))
	req, _ := http.NewRequest(http.MethodPost, "/tasks", reqBody)
	req.Header.Set("Content-Type", "application/json")

	respRecorder := httptest.NewRecorder()
	suite.router.ServeHTTP(respRecorder, req)

	suite.Equal(http.StatusCreated, respRecorder.Code)
	suite.mockTaskUseCase.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestGetAllTasksHandler() {
	expectedTasks := []domain.Task{
		{Title: "Task1", Description: "Description1"},
		{Title: "Task2", Description: "Description2"},
	}

	suite.mockTaskUseCase.On("GetAllTasks").Return(expectedTasks, nil)

	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	respRecorder := httptest.NewRecorder()

	suite.router.ServeHTTP(respRecorder, req)

	suite.Equal(http.StatusOK, respRecorder.Code)

	var responseTasks []domain.Task
	err := json.Unmarshal(respRecorder.Body.Bytes(), &responseTasks)
	suite.NoError(err)
	suite.Equal(expectedTasks, responseTasks)

	suite.mockTaskUseCase.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestGetTaskByIDHandler() {
	taskID := primitive.NewObjectID()
	task := domain.Task{
		ID:          taskID,
		Title:       "Task1",
		Description: "Description1",
	}

	suite.mockTaskUseCase.On("GetTaskByID", "taskID").Return(&task, nil)

	req, _ := http.NewRequest(http.MethodGet, "/tasks/taskID", nil)
	respRecorder := httptest.NewRecorder()

	suite.router.ServeHTTP(respRecorder, req)

	suite.Equal(http.StatusOK, respRecorder.Code)

	var responseTask domain.Task
	err := json.Unmarshal(respRecorder.Body.Bytes(), &responseTask)
	suite.NoError(err)
	suite.Equal(task, responseTask)

	suite.mockTaskUseCase.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestUpdateTaskHandler() {
	taskID := primitive.NewObjectID()
	task := domain.Task{
		ID:          taskID,
		Title:       "Updated Task",
		Description: "Updated description",
	}

	suite.mockTaskUseCase.On("UpdateTask", &task).Return(nil)

	reqBody := bytes.NewBuffer([]byte(`{"id": taskID, "title": "Updated Task", "description": "Updated description"}`))
	req, _ := http.NewRequest(http.MethodPut, "/tasks/taskID", reqBody)
	req.Header.Set("Content-Type", "application/json")

	respRecorder := httptest.NewRecorder()
	suite.router.ServeHTTP(respRecorder, req)

	suite.Equal(http.StatusOK, respRecorder.Code)
	suite.mockTaskUseCase.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestDeleteTaskHandler() {
	taskID := "1"
	suite.mockTaskUseCase.On("DeleteTask", taskID).Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/tasks/1", nil)
	respRecorder := httptest.NewRecorder()

	suite.router.ServeHTTP(respRecorder, req)

	suite.Equal(http.StatusOK, respRecorder.Code)
	suite.mockTaskUseCase.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestRegisterUserHandler() {
	user := domain.User{
		Username: "testuser",
		Password: "testpass",
	}

	suite.mockUserUseCase.On("RegisterUser", user.Username, user.Password).Return(&user, nil)

	reqBody := bytes.NewBuffer([]byte(`{"username": "testuser", "password": "testpass"}`))
	req, _ := http.NewRequest(http.MethodPost, "/register", reqBody)
	req.Header.Set("Content-Type", "application/json")

	respRecorder := httptest.NewRecorder()
	suite.router.ServeHTTP(respRecorder, req)

	suite.Equal(http.StatusCreated, respRecorder.Code)
	suite.mockUserUseCase.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestLoginUserHandler() {
	user := domain.User{
		Username: "testuser",
		Password: "testpass",
	}

	token := "mockToken"
	suite.mockUserUseCase.On("LoginUser", user.Username, user.Password).Return(token, nil)

	reqBody := bytes.NewBuffer([]byte(`{"username": "testuser", "password": "testpass"}`))
	req, _ := http.NewRequest(http.MethodPost, "/login", reqBody)
	req.Header.Set("Content-Type", "application/json")

	respRecorder := httptest.NewRecorder()
	suite.router.ServeHTTP(respRecorder, req)

	suite.Equal(http.StatusOK, respRecorder.Code)

	var response map[string]string
	err := json.Unmarshal(respRecorder.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Equal(token, response["token"])

	suite.mockUserUseCase.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestPromoteUserHandler() {
	userID := "1"
	suite.mockUserUseCase.On("PromoteUser", userID).Return(nil)

	req, _ := http.NewRequest(http.MethodPut, "/promote/1", nil)
	respRecorder := httptest.NewRecorder()

	suite.router.ServeHTTP(respRecorder, req)

	suite.Equal(http.StatusOK, respRecorder.Code)
	suite.mockUserUseCase.AssertExpectations(suite.T())
}

func TestControllerSuite(t *testing.T) {
	suite.Run(t, new(ControllerTestSuite))
}
