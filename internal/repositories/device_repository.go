package repositories

import (
	"context"
	"fmt"
	"github.com/Panitnun-6243/duckduck-server/internal/db"
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

//func IsDeviceAvailable(deviceCode string) bool {
//	var device models.DeviceRegistrationCode
//	err := db.GetDB().Collection("devices").FindOne(context.TODO(), bson.M{"code": deviceCode, "is_used": false}).Decode(&device)
//	return err == nil
//}
//
//func MarkDeviceAsUsed(deviceCode string) error {
//	_, err := db.GetDB().Collection("devices").UpdateOne(context.TODO(), bson.M{"code": deviceCode}, bson.M{"$set": bson.M{"is_used": true}})
//	return err
//}

func IsDeviceAvailable(deviceCode string) bool {
	var device models.DeviceRegistrationCode
	err := db.GetDB().Collection("devices").FindOne(context.TODO(), bson.M{"device_code": deviceCode, "is_used": false}).Decode(&device)
	if err == mongo.ErrNoDocuments {
		return false // No available device found with the provided code
	}
	if err != nil {
		log.Println(fmt.Sprintf("Error while fetching device by code: %v", err))
		return false
	}
	return true
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
