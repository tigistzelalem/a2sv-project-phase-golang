package data

import (
	"context"
	"errors"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskService struct {
	collection *mongo.Collection
}

func NewTaskService(client *mongo.Client) *TaskService {
	return &TaskService{
		collection: client.Database("task_manager").Collection("tasks"),
	}
}

// GetTasks retrieves all tasks from the database.
func (ts *TaskService) GetTasks() ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := ts.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

// GetTask retrieves a single task by its ID.
func (ts *TaskService) GetTask(id string) (*models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid task ID")
	}

	var task models.Task
	err = ts.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("task not found")
		}
		return nil, err
	}

	return &task, nil
}

// CreateTask adds a new task to the database.
func (ts *TaskService) CreateTask(task *models.Task) (*models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	task.ID = primitive.NewObjectID()

	_, err := ts.collection.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// UpdateTask updates an existing task in the database.
func (ts *TaskService) UpdateTask(id string, updatedTask *models.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task ID")
	}

	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"status":      updatedTask.Status,
			"updatedAt":   time.Now(),
		},
	}

	_, err = ts.collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}

	return nil
}

// DeleteTask removes a task from the database.
func (ts *TaskService) DeleteTask(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task ID")
	}

	_, err = ts.collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	return nil
}
