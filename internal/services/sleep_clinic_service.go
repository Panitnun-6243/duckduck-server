package services

import (
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"github.com/Panitnun-6243/duckduck-server/internal/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateDefaultSleepClinic(userID primitive.ObjectID) (*models.SleepClinic, error) {
	sleepClinic := &models.SleepClinic{
		UserID:                 userID,
		CurrentLullabySong:     "Twinkle twinkle little star",
		CurrentLullabySongPath: "https://storage.googleapis.com/duckduck-bucket/lullaby-song/noice/twinkle-twinkle-little-star.mp3",
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
		"sleep_stats":               updatedData.SleepStats,
		"current_lullaby_song":      updatedData.CurrentLullabySong,
		"current_lullaby_song_path": updatedData.CurrentLullabySongPath,
		"dim_light":                 updatedData.DimLight,
	}
	err := repositories.UpdateSleepClinicData(sleepClinicID, updateMap)

	return err
}

// AddCustomLullabySongService handles business logic for adding a custom lullaby song.
func AddCustomLullabySongService(userID primitive.ObjectID, songName, songPath, category string) error {
	song := models.LullabyDetail{Name: songName, Path: songPath, Category: category}
	return repositories.AddCustomLullabySong(userID, song)
}

// GetCustomLullabySongsService retrieves custom lullaby songs for the specified user.
func GetCustomLullabySongsService(userID primitive.ObjectID) ([]models.LullabyDetail, error) {
	return repositories.GetCustomLullabySongs(userID)
}

// GetPresetLullabySongsService retrieves all preset lullaby songs.
func GetPresetLullabySongsService() ([]models.PresetLullabySong, error) {
	return repositories.GetPresetLullabySongs()
}
