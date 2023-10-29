package repositories

import (
	"context"
	"github.com/Panitnun-6243/duckduck-server/database"
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
