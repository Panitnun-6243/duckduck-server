package repositories

import (
	"context"
	"github.com/Panitnun-6243/duckduck-server/db"
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func CreateAlarm(alarm *models.Alarm) (*models.Alarm, error) {
	_, err := db.GetDB().Collection("alarms").InsertOne(context.TODO(), alarm)
	if err != nil {
		log.Printf("Error while inserting alarm: %v", err)
		return nil, err
	}
	return alarm, nil
}

func FindAlarmByUserID(userID primitive.ObjectID) ([]*models.Alarm, error) {
	cursor, err := db.GetDB().Collection("alarms").Find(context.TODO(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var alarms []*models.Alarm
	if err = cursor.All(context.TODO(), &alarms); err != nil {
		return nil, err
	}
	return alarms, nil
}

func UpdateAlarm(alarmID primitive.ObjectID, updatedData bson.M) error {
	_, err := db.GetDB().Collection("alarms").UpdateOne(context.TODO(), bson.M{"_id": alarmID}, bson.M{"$set": updatedData})
	return err
}

func DeleteAlarm(alarmID primitive.ObjectID) error {
	_, err := db.GetDB().Collection("alarms").DeleteOne(context.TODO(), bson.M{"_id": alarmID})
	return err
}
