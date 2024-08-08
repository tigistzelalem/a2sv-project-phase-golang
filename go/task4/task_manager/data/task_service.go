package data

import (
	"errors"
	"task_manager/models"
)

var tasks = []models.Task{}
var id = 1

func GetAllTasks()[]models.Task  {
	return tasks	
}

func GetTaskById(id int)(models.Task, error)  {
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}

	return models.Task{}, errors.New("task not found")
	
}

func CreateTask(task models.Task) models.Task  {
	task.ID = id
	id ++
	tasks = append(tasks, task)
	return task
	
}

func UpdateTask(id int, updatedTask models.Task)(models.Task, error)  {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Title = updatedTask.Title
			tasks[i].Description = updatedTask.Description
			tasks[i].DueDate = updatedTask.DueDate
			tasks[i].Status = updatedTask.Status

			return tasks[i], nil
		}
	}

	return models.Task{}, errors.New("task not found")
	
}


func DeleteTask(id int) error  {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]... )
			return nil
		}
	}

	return errors.New("task not found")
	
}

