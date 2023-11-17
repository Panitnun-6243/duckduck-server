package services

import (
	"encoding/json"
	"fmt"
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"github.com/Panitnun-6243/duckduck-server/internal/repositories"
	"github.com/Panitnun-6243/duckduck-server/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
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

func TriggerAlarm(userID primitive.ObjectID, alarmID primitive.ObjectID) error {
	// Fetch the device code for the user
	deviceCode, err := repositories.FindDeviceCodeByUserID(userID)
	if err != nil {
		log.Printf("Error while fetching device code: %v", err)
		return err
	}

	// Proceed to publish MQTT event with the retrieved device code
	mqttTopic := fmt.Sprintf("%s/trigger-alarm", deviceCode)
	payload, _ := json.Marshal(map[string]string{"id": alarmID.Hex()})
	client := util.CreateMqttClient()
	util.Publish(client, mqttTopic, string(payload))

	return nil
}
