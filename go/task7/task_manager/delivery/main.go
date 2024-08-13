package main

import (
	"context"
	"log"
	"os"
	repositories "task-manager/Repositories"
	usecases "task-manager/Usecases"
	"task-manager/delivery/controllers"
	"task-manager/delivery/routers"
	"task-manager/infrastructure"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file", err)
	}

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI not set in .env file")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("failed to connect ot mongodb: %v", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("failed to ping mongodb")
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Failed to disconnect MongoDB: %v", err)
		}
	}()
	database := client.Database("task_manager")

	userRepo := repositories.NewUserRepository(database)
	taskRepo := repositories.NewTaskRepository(database)

	jwtService := infrastructure.NewJWTService()
	passwordService := infrastructure.NewPasswordService()

	userUsecase := usecases.NewUserUseCase(userRepo, passwordService, jwtService)
	taskUsecase := usecases.NewTaskUseCase(taskRepo)

	ctrl := controllers.NewController(userUsecase, taskUsecase)

	router := routers.SetUpRouter(ctrl, jwtService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port
	}

	// Start the server
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
