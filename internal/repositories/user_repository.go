package repositories

import (
	"context"
	"fmt"
	"github.com/Panitnun-6243/duckduck-server/internal/db"
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func CreateUser(user *models.User) (*models.User, error) {
	_, err := db.GetDB().Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		log.Println(fmt.Sprintf("Error while inserting user: %v", err))
		return nil, err
	}
	return user, nil
}

func FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := db.GetDB().Collection("users").FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		log.Println("No user found with the provided email")
		return nil, nil
	}
	if err != nil {
		log.Println(fmt.Sprintf("Error while fetching user by email: %v", err))
		return nil, err
	}
	return &user, nil
}
