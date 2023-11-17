package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/Panitnun-6243/duckduck-server/database"
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func FindDeviceByCode(deviceCode string) (*models.Device, error) {
	var device models.Device
	err := db.GetDB().Collection("devices").FindOne(context.TODO(), bson.M{"device_code": deviceCode}).Decode(&device)
	if err != nil {
		return nil, err
	}
	return &device, nil
}

func IsDeviceAvailable(deviceCode string) bool {
	device, err := FindDeviceByCode(deviceCode)
	return err == nil && !device.IsUsed
}

func MarkDeviceAsUsed(deviceCode string) error {
	_, err := db.GetDB().Collection("devices").UpdateOne(
		context.TODO(),
		bson.M{"device_code": deviceCode},
		bson.M{"$set": bson.M{"is_used": true}},
	)
	if err != nil {
		log.Println(fmt.Sprintf("Error while marking device as used: %v", err))
		return err
	}
	return nil
}

func BindUserToDevice(userID primitive.ObjectID, deviceCode string) error {
	_, err := db.GetDB().Collection("devices").UpdateOne(
		context.TODO(),
		bson.M{"device_code": deviceCode},
		bson.M{"$set": bson.M{"is_used": true, "user_id": userID}},
	)
	if err != nil {
		log.Println(fmt.Sprintf("Error while binding user to device: %v", err))
		return err
	}
	return nil
}

func FindUserByDeviceCode(deviceCode string) (*models.User, error) {
	var device models.Device
	err := db.GetDB().Collection("devices").FindOne(context.TODO(), bson.M{"device_code": deviceCode}).Decode(&device)
	if err != nil {
		return nil, err
	}
	if device.UserID.IsZero() {
		return nil, errors.New("no user bound to this device")
	}
	return FindUserByID(device.UserID.Hex())
}

// FindDeviceCodeByUserID fetches the device code associated with a user ID.
func FindDeviceCodeByUserID(userID primitive.ObjectID) (string, error) {
	var device models.Device
	err := db.GetDB().Collection("devices").FindOne(context.TODO(), bson.M{"user_id": userID}).Decode(&device)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("no device found for user")
		}
		log.Printf("Error while fetching device by user ID: %v", err)
		return "", err
	}
	return device.Code, nil
}
