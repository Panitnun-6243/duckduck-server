package repositories

import (
	"context"
	"fmt"
	"github.com/Panitnun-6243/duckduck-server/db"
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func CreateUser(user *models.User) (*models.User, error) {
	// Manual set user id to new object id
	user.ID = primitive.NewObjectID()
	_, err := db.GetDB().Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		log.Println(fmt.Sprintf("Error while inserting user: %v", err))
		return nil, err
	}
	return user, nil
}

// FindUserByEmail for check existing user
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

func FindUserByID(userID string) (*models.User, error) {
	var user models.User
	oid, _ := primitive.ObjectIDFromHex(userID)
	err := db.GetDB().Collection("users").FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUserDetails(userID, name, avatarURL string) (*models.User, error) {
	oid, _ := primitive.ObjectIDFromHex(userID)
	update := bson.M{
		"$set": bson.M{
			"name":       name,
			"avatar_url": avatarURL,
		},
	}
	_, err := db.GetDB().Collection("users").UpdateOne(context.TODO(), bson.M{"_id": oid}, update)
	if err != nil {
		return nil, err
	}
	return FindUserByID(userID)
}
