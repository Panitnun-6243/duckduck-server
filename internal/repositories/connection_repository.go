package repositories

import (
	"context"
	"github.com/Panitnun-6243/duckduck-server/db"
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindConnectionStatusByUserID(userID primitive.ObjectID) (bool, error) {
	var connectionStatus models.LightControl
	err := db.GetDB().Collection("light_controls").FindOne(context.TODO(), bson.M{"user_id": userID}, options.FindOne().SetProjection(bson.M{"connected": 1})).Decode(&connectionStatus)
	return connectionStatus.Connected, err
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

func UpdateConnectionStatus(connID primitive.ObjectID, connected bool) error {
	_, err := db.GetDB().Collection("light_controls").UpdateOne(context.TODO(), bson.M{"_id": connID}, bson.M{"$set": bson.M{"connected": connected}})
	return err
}
