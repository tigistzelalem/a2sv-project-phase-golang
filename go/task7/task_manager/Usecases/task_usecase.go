package usecases

import (
	repositories "task-manager/Repositories"
	"task-manager/domain"
)

type TaskUseCase interface {
	CreateTask(task *domain.Task) error
	GetAllTasks() ([]domain.Task, error)
	GetTaskByID(id string)(*domain.Task, error)
	UpdateTask(task *domain.Task) error
	DeleteTask (id string) error
}

type taskUseCase struct {
	taskRepo repositories.TaskRepository
}

func NewTaskUseCase(taskRepo repositories.TaskRepository) TaskUseCase {
	return &taskUseCase{
		taskRepo: taskRepo,
	}
}


func (uc *taskUseCase) CreateTask(task *domain.Task) error {
	return uc.taskRepo.CreateTask(task)
}

func (uc *taskUseCase) GetAllTasks() ([]domain.Task, error)  {
	return uc.taskRepo.GetAllTasks()
	
}

func (uc *taskUseCase) GetTaskByID(id string) (*domain.Task, error) {
	return uc.taskRepo.GetTaskByID(id)
}

func (uc *taskUseCase) UpdateTask(task *domain.Task) error {
	return uc.taskRepo.UpdateTask(task)
}

func (uc *taskUseCase) DeleteTask(id string) error {
	return uc.taskRepo.DeleteTask(id)
}