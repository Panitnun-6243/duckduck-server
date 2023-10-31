package services

import (
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"github.com/Panitnun-6243/duckduck-server/internal/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCctLightByUser(userID primitive.ObjectID) (*models.LightControl, error) {
	return repositories.FindCctLightByUserID(userID)
}

func GetCctLightByIDAndUserID(cctID, userID primitive.ObjectID) (*models.LightControl, error) {
	return repositories.FindCctLightByIDAndUserID(cctID, userID)
}

func UpdateUserCctLight(cctID primitive.ObjectID, brightness float64, temp int) error {
	return repositories.UpdateCctLight(cctID, brightness, temp)
}
