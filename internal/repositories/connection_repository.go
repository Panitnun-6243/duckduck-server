package repositories

import (
	"context"
	"github.com/Panitnun-6243/duckduck-server/db"
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindConnectionStatusByUserID(userID primitive.ObjectID) (*models.LightControl, error) {
	var connectionStatus *models.LightControl
	err := db.GetDB().Collection("light_controls").FindOne(context.TODO(), bson.M{"user_id": userID}).Decode(&connectionStatus)
	return connectionStatus, err
}

func FindConnectionStatusByIDAndUserID(connID, userID primitive.ObjectID) (*models.LightControl, error) {
	var connectionStatus *models.LightControl
	filter := bson.M{
		"_id":     connID,
		"user_id": userID,
	}
	err := db.GetDB().Collection("light_controls").FindOne(context.TODO(), filter).Decode(&connectionStatus)
	return connectionStatus, err
}

func UpdateConnectionStatus(connID primitive.ObjectID, updatedData bson.M) error {
	_, err := db.GetDB().Collection("light_controls").UpdateOne(context.TODO(), bson.M{"_id": connID}, bson.M{"$set": updatedData})
	return err
}
