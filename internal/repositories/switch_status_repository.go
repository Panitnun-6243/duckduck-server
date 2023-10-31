package repositories

import (
	"context"
	db "github.com/Panitnun-6243/duckduck-server/database"
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindSwitchStatusByUserID(userID primitive.ObjectID) (bool, error) {
	var switchStatus models.LightControl
	err := db.GetDB().Collection("light_controls").FindOne(context.TODO(), bson.M{"user_id": userID}, options.FindOne().SetProjection(bson.M{"on": 1})).Decode(&switchStatus)
	return switchStatus.SwitchStatus, err
}

func FindSwitchStatusByIDAndUserID(switchID, userID primitive.ObjectID) (*models.LightControl, error) {
	var switchStatus *models.LightControl
	filter := bson.M{
		"_id":     switchID,
		"user_id": userID,
	}
	err := db.GetDB().Collection("light_controls").FindOne(context.TODO(), filter).Decode(&switchStatus)
	return switchStatus, err
}

func UpdateSwitchStatus(switchID primitive.ObjectID, on bool) error {
	_, err := db.GetDB().Collection("light_controls").UpdateOne(context.TODO(), bson.M{"_id": switchID}, bson.M{"$set": bson.M{"on": on}})
	return err
}
