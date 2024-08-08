package data

import (
	"context"
	"log"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var taskCollection *mongo.Collection

func INitMongoDB() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://makgg015:MflcGi2k4MOQehad@task.gf9ah.mongodb.net/?retryWrites=true&w=majority&appName=task")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	taskCollection = client.Database("task_manager_db").Collection("tasks")
}



func GetAllTasks() ([]models.Task, error) {
    var tasks []models.Task
    cursor, err := taskCollection.Find(context.TODO(), bson.D{{}})
    if err != nil {
        return nil, err
    }

    if err = cursor.All(context.TODO(), &tasks); err != nil {
        return nil, err
    }

    return tasks, nil
}

func GetTaskByID(id string) (models.Task, error) {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return models.Task{}, err
    }

    var task models.Task
    err = taskCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&task)
    if err != nil {
        return models.Task{}, err
    }

    return task, nil
}

func CreateTask(task models.Task)(*mongo.InsertOneResult, error)  {
	task.ID = primitive.NewObjectID()
	return taskCollection.InsertOne(context.TODO(), task)
	
}

func UpdateTask(id string, task models.Task) error  {
	 objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return  nil
    }

    _, err = taskCollection.UpdateOne(
        context.TODO(),
        bson.M{"_id": objID},
        bson.D{
            {"$set", bson.D{
                {"title", task.Title},
                {"description", task.Description},
                {"due_date", task.DueDate},
                {"status", task.Status},
            }},
        },
    )
    return err
	
}


func DeleteTask(id string) error  {
	 objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    _, err = taskCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
    return err
	
}

