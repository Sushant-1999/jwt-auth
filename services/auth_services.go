package services

import (
	"context"
	"errors"
	"jwt-auth-service/config"
	"jwt-auth-service/models"
	"jwt-auth-service/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection

// func init() {
// 	userCollection = config.GetDBCollection("users")
// }

func CreateUser(user *models.User) error {
	userCollection = config.GetDBCollection("users")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	user.ID = primitive.NewObjectID()
	_, err = userCollection.InsertOne(context.Background(), user)
	return err
}

func AuthenticateUser(credentials *models.User) (string, error) {
	userCollection = config.GetDBCollection("users")
	var user models.User
	err := userCollection.FindOne(context.Background(), bson.M{"username": credentials.Username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("invalid username or password")
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := utils.GenerateToken(user.ID.Hex())
	if err != nil {
		return "", err
	}

	return token, nil
}
