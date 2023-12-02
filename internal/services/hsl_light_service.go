package services

import (
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"github.com/Panitnun-6243/duckduck-server/internal/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetHslLightByUser(userID primitive.ObjectID) (models.Hsl, error) {
	return repositories.FindHslLightByUserID(userID)
}

func GetHslLightByIDAndUserID(hslID, userID primitive.ObjectID) (*models.LightControl, error) {
	return repositories.FindHslLightByIDAndUserID(hslID, userID)
}

func UpdateUserHslLight(hslID primitive.ObjectID, hslLight models.Hsl) error {
	return repositories.UpdateHslLight(hslID, hslLight)
}

// UpdateUserHslLightColorMode Update only color_mode
func UpdateUserHslLightColorMode(hslID primitive.ObjectID, colorMode string) error {
	return repositories.UpdateLightControlColorMode(hslID, colorMode)
}
