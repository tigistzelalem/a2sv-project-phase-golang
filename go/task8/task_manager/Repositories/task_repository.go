package repositories

import (
	"context"
	"task-manager/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository interface {
	CreateTask(task *domain.Task) error
	GetAllTasks() ([]domain.Task, error)
	GetTaskByID(id string) (*domain.Task, error)
	UpdateTask(task *domain.Task) error
	DeleteTask(id string) error
}

type taskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository(db *mongo.Database) TaskRepository {
	return &taskRepository{
		collection: db.Collection("tasks"),
	}
}


func (repo *taskRepository) CreateTask(task *domain.Task) error {
	_, err := repo.collection.InsertOne(context.Background(), task)
	return err
}

func (repo *taskRepository) GetAllTasks() ([]domain.Task, error) {
	cursor, err := repo.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var tasks []domain.Task
	for cursor.Next(context.Background()) {
		var task domain.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (repo *taskRepository) GetTaskByID(id string) (*domain.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var task domain.Task
	err = repo.collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&task)
	return &task, err
}

func (repo *taskRepository) UpdateTask(task *domain.Task) error {
	_, err := repo.collection.UpdateOne(context.Background(), bson.M{"_id": task.ID}, bson.M{"$set": task})
	return err
}

func (repo *taskRepository) DeleteTask(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = repo.collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	return err
}