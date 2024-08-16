package controllers

import (
	"net/http"
	usecases "task-manager/Usecases"
	"task-manager/domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Controller struct {
	userUC usecases.UserUseCase
	taskUC usecases.TaskUseCase
}

func NewController(userUsecase usecases.UserUseCase, taskUsecase usecases.TaskUseCase) *Controller {
	return &Controller{
		userUC: userUsecase,
		taskUC: taskUsecase,
	}
}

func (ctrl *Controller) RegisterUser(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	}

	user, err := ctrl.userUC.RegisterUser(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	c.JSON(http.StatusCreated, user)

}

func (ctrl *Controller) LoginUser(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
	}

	token, err := ctrl.userUC.LoginUser(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func (ctrl *Controller) PromoteUser(c *gin.Context) {
	userID := c.Param("id")
	if err := ctrl.userUC.PromoteUser(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user promoted"})

}

// task controller

func (ctrl *Controller) CreateTask(c *gin.Context) {
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
	}

	if err := ctrl.taskUC.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)

}

func (ctrl *Controller) GetAllTasks(c *gin.Context) {
	tasks, err := ctrl.taskUC.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internak server error"})
		return
	}

	c.JSON(http.StatusOK, tasks)

}

func (ctrl *Controller) GetTaskByID(c *gin.Context) {
	taskID := c.Param("id")
	task, err := ctrl.taskUC.GetTaskByID(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)

}

func (ctrl *Controller) UpdateTask(c *gin.Context) {
	taskID := c.Param("id")
	var task domain.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	task.ID, _ = primitive.ObjectIDFromHex(taskID)
	if err := ctrl.taskUC.UpdateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)

}

func (ctrl *Controller) DeleteTask(c *gin.Context) {
	taskID := c.Param("id")
	if err := ctrl.taskUC.DeleteTask(taskID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})

}
