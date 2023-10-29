package services

import (
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"github.com/Panitnun-6243/duckduck-server/internal/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateDefaultSleepClinic(userID primitive.ObjectID) (*models.SleepClinic, error) {
	sleepClinic := &models.SleepClinic{
		UserID:                userID,
		CurrentLullabySong:    "Twinkle twinkle little star",
		CustomLullabySongPath: "",
		DimLight: models.DimLight{
			IsActive: false,
			Duration: 5,
		},
	}
	return repositories.CreateDefaultSleepClinicData(sleepClinic)
}

func GetSleepClinicByUser(userID primitive.ObjectID) (*models.SleepClinic, error) {
	return repositories.FindSleepClinicByUserID(userID)
}

func UpdateUserSleepClinic(sleepClinicID primitive.ObjectID, updatedData *models.SleepClinic) error {
	updateMap := bson.M{
		"sleep_stats":              updatedData.SleepStats,
		"current_lullaby_song":     updatedData.CurrentLullabySong,
		"custom_lullaby_song_path": updatedData.CustomLullabySongPath,
		"dim_light":                updatedData.DimLight,
	}
	err := repositories.UpdateSleepClinicData(sleepClinicID, updateMap)
	//utils.Publish("sleep_clinic/update", "Updated Sleep Clinic Data")
	return err
}
