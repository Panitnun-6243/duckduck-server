package services

import (
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"github.com/Panitnun-6243/duckduck-server/internal/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func CreateDefaultLightControl(user primitive.ObjectID) (*models.LightControl, error) {
	defaultLightControlData := &models.LightControl{
		UserID:               user,
		Connected:            false,
		SwitchStatus:         "off",
		BrightnessPercentage: 51.0,
		ColorMode:            "cct",
		CctTemp:              3000,
		RgbColor:             "0xfff2DB6F",
	}
	defaultLightControlData.CreatedAt = time.Now()
	defaultLightControlData.UpdatedAt = time.Now()

	return repositories.CreateLightControl(defaultLightControlData)
}

func GetLightControlByUser(userID primitive.ObjectID) (*models.LightControl, error) {
	return repositories.FindLightControlByUserID(userID)
}

func GetLightControlByIDAndUserID(controlID, userID primitive.ObjectID) (*models.LightControl, error) {
	return repositories.FindLightControlByIDAndUserID(controlID, userID)
}

func UpdateUserLightControl(controlID primitive.ObjectID, updatedData *models.LightControl) error {
	updatedData.UpdatedAt = time.Now()

	updateMap := bson.M{
		"switch_status":         updatedData.SwitchStatus,
		"connected":             updatedData.Connected,
		"brightness_percentage": updatedData.BrightnessPercentage,
		"color_mode":            updatedData.ColorMode,
		"cct_temp":              updatedData.CctTemp,
		"rgb_color":             updatedData.RgbColor,
		"updated_at":            updatedData.UpdatedAt,
	}

	err := repositories.UpdateLightControl(controlID, updateMap)
	return err
}
