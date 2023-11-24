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

func CreateDefaultSleepClinicData(sleepClinic *models.SleepClinic) (*models.SleepClinic, error) {
	_, err := db.GetDB().Collection("sleep_clinics").InsertOne(context.TODO(), sleepClinic)
	if err != nil {
		log.Printf("Error while inserting default sleep clinic data: %v", err)
		return nil, err
	}
	return sleepClinic, nil
}

func FindSleepClinicByUserID(userID primitive.ObjectID) (*models.SleepClinic, error) {
	var sleepClinic *models.SleepClinic
	err := db.GetDB().Collection("sleep_clinics").FindOne(context.TODO(), bson.M{"user_id": userID}).Decode(&sleepClinic)
	return sleepClinic, err
}

func UpdateSleepClinicData(sleepClinicID primitive.ObjectID, updatedData bson.M) error {
	_, err := db.GetDB().Collection("sleep_clinics").UpdateOne(context.TODO(), bson.M{"_id": sleepClinicID}, bson.M{"$set": updatedData})
	return err
}

// AddCustomLullabySong adds a new custom lullaby song to the user's custom lullaby songs list.
func AddCustomLullabySong(userID primitive.ObjectID, song models.LullabyDetail) error {
	collection := db.GetDB().Collection("custom_lullaby_songs")
	// Check for duplicate song name
	var existingData models.CustomLullabySong
	err := collection.FindOne(context.TODO(), bson.M{
		"user_id":    userID,
		"songs.name": song.Name,
	}).Decode(&existingData)
	if err == nil {
		return errors.New("song name already exists")
	}
	_, err = collection.UpdateOne(
		context.TODO(),
		bson.M{"user_id": userID},
		bson.M{"$push": bson.M{"songs": song}},
		options.Update().SetUpsert(true),
	)
	return err
}

// GetCustomLullabySongs retrieves custom lullaby songs for a given user.
func GetCustomLullabySongs(userID primitive.ObjectID) ([]models.LullabyDetail, error) {
	var customLullabySong models.CustomLullabySong
	collection := db.GetDB().Collection("custom_lullaby_songs")
	err := collection.FindOne(context.TODO(), bson.M{"user_id": userID}).Decode(&customLullabySong)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []models.LullabyDetail{}, nil
		}
		return nil, err
	}
	return customLullabySong.Songs, nil
}

// GetPresetLullabySongs returns all preset lullaby songs.
func GetPresetLullabySongs() ([]models.PresetLullabySong, error) {
	var songs []models.PresetLullabySong
	collection := db.GetDB().Collection("preset_lullaby_songs")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &songs); err != nil {
		return nil, err
	}
	return songs, nil
}
