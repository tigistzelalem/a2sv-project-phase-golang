package controllers

import (
	"net/http"
	"strconv"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {

	 tasks := data.GetAllTasks()
	 c.JSON(http.StatusOK, tasks)
}

func GetTask(c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
	}

	task, err := data.GetTaskById(id)
	if err != nil {
		  c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
	}

	 c.JSON(http.StatusOK, task)
	
}

func CreateTask(c *gin.Context)  {
	
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	createdTask := data.CreateTask(newTask)
	c.JSON(http.StatusCreated, createdTask)
}


func UpdateTask(c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
	}
	
	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	task, err := data.UpdateTask(id, updatedTask)
	if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }

    err = data.DeleteTask(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.Status(http.StatusNoContent)
}