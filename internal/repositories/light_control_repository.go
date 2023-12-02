package repositories

import (
	"context"
	"github.com/Panitnun-6243/duckduck-server/database"
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func CreateLightControl(lightControl *models.LightControl) (*models.LightControl, error) {
	_, err := db.GetDB().Collection("light_controls").InsertOne(context.TODO(), lightControl)
	if err != nil {
		log.Printf("Error while inserting light control: %v", err)
		return nil, err
	}
	return lightControl, nil
}

func FindLightControlByUserID(userID primitive.ObjectID) (*models.LightControl, error) {
	var lightControl *models.LightControl
	err := db.GetDB().Collection("light_controls").FindOne(context.TODO(), bson.M{"user_id": userID}).Decode(&lightControl)
	return lightControl, err
}

func FindLightControlByIDAndUserID(controlID, userID primitive.ObjectID) (*models.LightControl, error) {
	var lightControl *models.LightControl
	filter := bson.M{
		"_id":     controlID,
		"user_id": userID,
	}
	err := db.GetDB().Collection("light_controls").FindOne(context.TODO(), filter).Decode(&lightControl)
	return lightControl, err
}

func UpdateLightControl(controlID primitive.ObjectID, updatedData bson.M) error {
	_, err := db.GetDB().Collection("light_controls").UpdateOne(context.TODO(), bson.M{"_id": controlID}, bson.M{"$set": updatedData})
	return err
}

// UpdateLightControlColorMode Update only color_mode
func UpdateLightControlColorMode(controlID primitive.ObjectID, colorMode string) error {
	_, err := db.GetDB().Collection("light_controls").UpdateOne(context.TODO(), bson.M{"_id": controlID}, bson.M{"$set": bson.M{
		"color_mode": colorMode,
	}})
	return err
}
