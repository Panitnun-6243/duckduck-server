package services

import (
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"github.com/Panitnun-6243/duckduck-server/internal/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetSwitchStatusByUser(userID primitive.ObjectID) (bool, error) {
	return repositories.FindSwitchStatusByUserID(userID)
}

func GetSwitchStatusByIDAndUserID(switchID, userID primitive.ObjectID) (*models.LightControl, error) {
	return repositories.FindSwitchStatusByIDAndUserID(switchID, userID)
}

func UpdateUserSwitchStatus(switchID primitive.ObjectID, on bool) error {
	return repositories.UpdateSwitchStatus(switchID, on)
}
