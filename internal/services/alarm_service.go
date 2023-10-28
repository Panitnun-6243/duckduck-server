package services

import (
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"github.com/Panitnun-6243/duckduck-server/internal/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func CreateNewAlarm(userID primitive.ObjectID, alarmData *models.Alarm) (*models.Alarm, error) {
	alarmData.UserID = userID
	alarmData.CreatedAt = time.Now()
	alarmData.UpdatedAt = time.Now()

	return repositories.CreateAlarm(alarmData)
}

func GetAlarmsByUser(userID primitive.ObjectID) ([]*models.Alarm, error) {
	return repositories.FindAlarmByUserID(userID)
}

func UpdateUserAlarm(userID, alarmID primitive.ObjectID, updatedData *models.Alarm) error {
	updatedData.UpdatedAt = time.Now()
	updatedData.IsActive.DateActive = time.Now()
	_, err := repositories.UpdateAlarm(alarmID, updatedData)
	return err
}

func RemoveAlarm(alarmID primitive.ObjectID) error {
	return repositories.DeleteAlarm(alarmID)
}
