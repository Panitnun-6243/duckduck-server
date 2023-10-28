package services

import (
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"github.com/Panitnun-6243/duckduck-server/internal/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func CreateNewAlarm(alarmData *models.Alarm) (*models.Alarm, error) {
	alarmData.IsActive.DateActive = time.Now()
	alarmData.CreatedAt = time.Now()
	alarmData.UpdatedAt = time.Now()

	return repositories.CreateAlarm(alarmData)
}

func GetAlarmsByUser(userID primitive.ObjectID) ([]*models.Alarm, error) {
	return repositories.FindAlarmByUserID(userID)
}

func UpdateUserAlarm(alarmID primitive.ObjectID, updatedData *models.Alarm) error {
	updatedData.UpdatedAt = time.Now()
	updatedData.IsActive.DateActive = time.Now()
	// Convert updatedData into a map for the update operation
	updateMap := bson.M{
		"bed_time":                 updatedData.BedTime,
		"wake_up_time":             updatedData.WakeUpTime,
		"description":              updatedData.Description,
		"is_active":                updatedData.IsActive,
		"repeat_days":              updatedData.RepeatDays,
		"sunrise":                  updatedData.Sunrise,
		"current_wakeup_sound":     updatedData.CurrentWakeupSound,
		"custom_wakeup_sound_path": updatedData.CustomWakeupSoundPath,
		"volume":                   updatedData.Volume,
		"snooze_time":              updatedData.SnoozeTime,
		"updated_at":               updatedData.UpdatedAt,
	}

	err := repositories.UpdateAlarm(alarmID, updateMap)
	return err
}

func RemoveAlarm(alarmID primitive.ObjectID) error {
	return repositories.DeleteAlarm(alarmID)
}
