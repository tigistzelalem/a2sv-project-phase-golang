package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Controller struct {
	userService *data.UserService
	taskService *data.TaskService
}

func NewController(userService *data.UserService, taskService *data.TaskService) *Controller {
	return &Controller{
		userService: userService,
		taskService: taskService,
	}
}

// User-related handlers

func (c *Controller) RegisterUser(ctx *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.userService.RegisterUser(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (c *Controller) LoginUser(ctx *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.userService.Login(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (c *Controller) PromoteUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = c.userService.PromoteUser(objID.Hex())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User promoted"})
}

// Task-related handlers

func (c *Controller) GetTasks(ctx *gin.Context) {
	tasks, err := c.taskService.GetTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (c *Controller) GetTask(ctx *gin.Context) {
	id := ctx.Param("id")
	task, err := c.taskService.GetTask(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (c *Controller) CreateTask(ctx *gin.Context) {
	var task models.Task
	if err := ctx.BindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.taskService.CreateTask(&task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *Controller) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var task models.Task
	if err := ctx.BindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.taskService.UpdateTask(id, &task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Task updated successfully"})
}

func (c *Controller) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.taskService.DeleteTask(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Task deleted successfully"})
}
