package repositories

import (
	"context"
	"errors"
	"github.com/Panitnun-6243/duckduck-server/database"
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func CreateAlarm(alarm *models.Alarm) (*models.Alarm, error) {
	// Manual set user id to new object id
	alarm.ID = primitive.NewObjectID()
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

func FindAlarmByAlarmIDAndUserID(alarmID, userID primitive.ObjectID) (*models.Alarm, error) {
	var alarm *models.Alarm
	err := db.GetDB().Collection("alarms").FindOne(context.TODO(), bson.M{"_id": alarmID, "user_id": userID}).Decode(&alarm)
	if err != nil {
		return nil, err
	}
	return alarm, nil
}

func UpdateAlarm(alarmID primitive.ObjectID, updatedData bson.M) error {
	_, err := db.GetDB().Collection("alarms").UpdateOne(context.TODO(), bson.M{"_id": alarmID}, bson.M{"$set": updatedData})
	return err
}

func DeleteAlarm(alarmID primitive.ObjectID) error {
	_, err := db.GetDB().Collection("alarms").DeleteOne(context.TODO(), bson.M{"_id": alarmID})
	return err
}

// AddCustomAlarmSound adds a new custom alarm sound to the user's custom alarm sounds list.
func AddCustomAlarmSound(userID primitive.ObjectID, sound models.SoundDetail) error {
	collection := db.GetDB().Collection("custom_alarm_sounds")

	// Check for duplicate sound name
	var existingData models.CustomAlarmSound
	err := collection.FindOne(context.TODO(), bson.M{
		"user_id":     userID,
		"sounds.name": sound.Name,
	}).Decode(&existingData)
	if err == nil {
		return errors.New("sound name already exists")
	}

	_, err = collection.UpdateOne(
		context.TODO(),
		bson.M{"user_id": userID},
		bson.M{"$push": bson.M{"sounds": sound}},
		options.Update().SetUpsert(true),
	)
	return err
}

// GetCustomAlarmSounds retrieves custom alarm sounds for a given user.
func GetCustomAlarmSounds(userID primitive.ObjectID) ([]models.SoundDetail, error) {
	var customAlarmSound models.CustomAlarmSound
	collection := db.GetDB().Collection("custom_alarm_sounds")
	err := collection.FindOne(context.TODO(), bson.M{"user_id": userID}).Decode(&customAlarmSound)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Return an empty slice instead of an error when no documents are found
			return []models.SoundDetail{}, nil
		}
		return nil, err
	}
	return customAlarmSound.Sounds, nil
}

// GetPresetAlarmSounds returns all preset alarm sounds.
func GetPresetAlarmSounds() ([]models.PresetAlarmSound, error) {
	var sounds []models.PresetAlarmSound
	collection := db.GetDB().Collection("preset_alarm_sounds")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &sounds); err != nil {
		return nil, err
	}
	return sounds, nil
}
