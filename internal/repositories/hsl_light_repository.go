package repositories

import (
	"context"
	db "github.com/Panitnun-6243/duckduck-server/database"
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindHslLightByUserID(userID primitive.ObjectID) (models.Hsl, error) {
	var hslLight models.LightControl
	err := db.GetDB().Collection("light_controls").FindOne(context.TODO(), bson.M{"user_id": userID}, options.FindOne().SetProjection(bson.M{"hsl_color": 1})).Decode(&hslLight)
	return hslLight.HslColor, err
}

func FindHslLightByIDAndUserID(hslID, userID primitive.ObjectID) (*models.LightControl, error) {
	var hslLight *models.LightControl
	filter := bson.M{
		"_id":     hslID,
		"user_id": userID,
	}
	err := db.GetDB().Collection("light_controls").FindOne(context.TODO(), filter).Decode(&hslLight)
	return hslLight, err
}

func UpdateHslLight(hslID primitive.ObjectID, hslLight models.Hsl) error {
	_, err := db.GetDB().Collection("light_controls").UpdateOne(context.TODO(), bson.M{"_id": hslID}, bson.M{"$set": bson.M{"hsl_color": hslLight}})
	return err
}
