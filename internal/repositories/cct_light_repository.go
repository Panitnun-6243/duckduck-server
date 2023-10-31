package repositories

import (
	"context"
	db "github.com/Panitnun-6243/duckduck-server/database"
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindCctLightByUserID(userID primitive.ObjectID) (*models.LightControl, error) {
	var cctLight *models.LightControl
	err := db.GetDB().Collection("light_controls").FindOne(context.TODO(), bson.M{"user_id": userID}, options.FindOne().SetProjection(bson.M{"brightness": 1, "temp": 1})).Decode(&cctLight)
	return cctLight, err
}

func FindCctLightByIDAndUserID(cctID, userID primitive.ObjectID) (*models.LightControl, error) {
	var cctLight *models.LightControl
	filter := bson.M{
		"_id":     cctID,
		"user_id": userID,
	}
	err := db.GetDB().Collection("light_controls").FindOne(context.TODO(), filter).Decode(&cctLight)
	return cctLight, err
}

func UpdateCctLight(cctID primitive.ObjectID, brightness float64, temp int) error {
	_, err := db.GetDB().Collection("light_controls").UpdateOne(context.TODO(), bson.M{"_id": cctID}, bson.M{"$set": bson.M{"brightness": brightness, "temp": temp}})
	return err
}
