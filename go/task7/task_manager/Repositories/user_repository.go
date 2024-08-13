package repositories

import (
	"context"
	"task-manager/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	RegisterUser(user *domain.User) error
	LoginUser(username string) (*domain.User, error)
	PromoteUser(userID string) error
}


type userRepository struct {
	collection *mongo.Collection
}


func NewUserRepository (db *mongo.Database) UserRepository {
	return &userRepository {
		collection: db.Collection("users"),
	}
}

func (repo *userRepository) RegisterUser(user *domain.User) error {
	var existingUser domain.User
	err := repo.collection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&existingUser)
	if err == nil {
		return mongo.ErrClientDisconnected 
	}

	count, err := repo.collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	if count == 0 {
		user.Role = "admin"
	}

	_, err = repo.collection.InsertOne(context.Background(), user)
	return err
}

func (repo *userRepository) LoginUser(username string) (*domain.User, error) {
	var user domain.User
	err := repo.collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	return &user, err
}

func (repo *userRepository) PromoteUser(userID string) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	_, err = repo.collection.UpdateByID(context.Background(), objID, bson.M{"$set": bson.M{"role": "admin"}})
	return err
}
