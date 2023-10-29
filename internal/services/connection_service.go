package services

import (
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"github.com/Panitnun-6243/duckduck-server/internal/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetConnectionStatusByUser(userID primitive.ObjectID) (*models.LightControl, error) {
	return repositories.FindConnectionStatusByUserID(userID)
}

func GetConnectionStatusByIDAndUserID(connID, userID primitive.ObjectID) (*models.LightControl, error) {
	return repositories.FindConnectionStatusByIDAndUserID(connID, userID)
}

func UpdateUserConnectionStatus(connID primitive.ObjectID, updatedData *models.LightControl) error {

	updateMap := bson.M{
		"connected": updatedData.Connected,
	}

	err := repositories.UpdateConnectionStatus(connID, updateMap)
	return err
}
