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
		"current_alarm_sound":      updatedData.CurrentAlarmSound,
		"current_alarm_sound_path": updatedData.CurrentAlarmSoundPath,
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

	// Calculate sleep duration and update Sleep Clinic data
	updateSleepDuration(userID, alarmID)

	return nil
}

func updateSleepDuration(userID, alarmID primitive.ObjectID) {
	alarm, _ := repositories.FindAlarmByAlarmIDAndUserID(alarmID, userID)
	if alarm != nil {
		// Construct time.Time objects from BedTime and WakeUpTime
		now := time.Now()
		bedTime := time.Date(now.Year(), now.Month(), now.Day(), alarm.BedTime.Hours, alarm.BedTime.Minutes, 0, 0, now.Location())
		wakeUpTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, now.Location())

		// If wake-up time is before bedtime, it means wake up is on the next day
		if wakeUpTime.Before(bedTime) {
			wakeUpTime = wakeUpTime.AddDate(0, 0, 1)
		}

		// Calculate duration in hours
		duration := wakeUpTime.Sub(bedTime).Hours()

		// Update Sleep Clinic Data
		sleepClinic, _ := repositories.FindSleepClinicByUserID(userID)
		if sleepClinic != nil {
			sleepStat := models.SleepStat{
				Date:               now.Format("2006-01-02"),
				SleepDurationHours: duration,
			}
			sleepClinic.SleepStats = append(sleepClinic.SleepStats, sleepStat)
			_ = repositories.UpdateSleepClinicData(sleepClinic.ID, bson.M{"sleep_stats": sleepClinic.SleepStats})
		}
	}
}

// AddCustomAlarmSoundService handles business logic for adding a custom alarm sound.
func AddCustomAlarmSoundService(userID primitive.ObjectID, soundName, soundPath string) error {
	sound := models.SoundDetail{Name: soundName, Path: soundPath}
	return repositories.AddCustomAlarmSound(userID, sound)
}

// GetCustomAlarmSoundsService retrieves custom alarm sounds for the specified user.
func GetCustomAlarmSoundsService(userID primitive.ObjectID) ([]models.SoundDetail, error) {
	return repositories.GetCustomAlarmSounds(userID)
}

// GetPresetAlarmSoundsService retrieves all preset alarm sounds.
func GetPresetAlarmSoundsService() ([]models.PresetAlarmSound, error) {
	return repositories.GetPresetAlarmSounds()
}
