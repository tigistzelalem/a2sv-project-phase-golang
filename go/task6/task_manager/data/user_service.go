package data

import (
    "context"
    "errors"
    "time"
    "task_manager/models"
    "github.com/dgrijalva/jwt-go"
    "golang.org/x/crypto/bcrypt"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
    collection *mongo.Collection
    jwtSecret  []byte
}

func NewUserService(client *mongo.Client, jwtSecret string) *UserService {
    return &UserService{
        collection: client.Database("task_manager").Collection("users"),
        jwtSecret:  []byte(jwtSecret),
    }
}

func (us *UserService) RegisterUser(username, password string) (*models.User, error) {
    var existingUser models.User
    err := us.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&existingUser)
    if err == nil {
        return nil, errors.New("username already exists")
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    user := models.User{
        Username: username,
        Password: string(hashedPassword),
        Role:     "user", // Default role
    }

    // Check if this is the first user
    count, err := us.collection.CountDocuments(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }

    if count == 0 {
        // Set the first user as an admin
        user.Role = "admin"
    }

    _, err = us.collection.InsertOne(context.Background(), user)
    if err != nil {
        return nil, err
    }

    return &user, nil
}

func (us *UserService) Login(username, password string) (string, error) {
    var user models.User
    err := us.collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
    if err != nil {
        return "", errors.New("invalid credentials")
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return "", errors.New("invalid credentials")
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id":       user.ID.Hex(),
        "username": user.Username,
        "role":     user.Role,
        "exp":      time.Now().Add(time.Hour * 24).Unix(),
    })

    signedToken, err := token.SignedString(us.jwtSecret)
    if err != nil {
        return "", err
    }

    return signedToken, nil
}

func (us *UserService) PromoteUser(userID string) error {
    // Function to promote user to admin
    _, err := us.collection.UpdateByID(context.Background(), userID, bson.M{"$set": bson.M{"role": "admin"}})
    return err
}
