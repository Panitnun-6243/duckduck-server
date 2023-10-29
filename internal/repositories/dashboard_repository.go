package repositories

import (
	"context"
	"github.com/Panitnun-6243/duckduck-server/db"
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func CreateDashboardConfig(config *models.DashboardConfig) (*models.DashboardConfig, error) {
	_, err := db.GetDB().Collection("dashboard_configs").InsertOne(context.TODO(), config)
	if err != nil {
		log.Printf("Error while inserting dashboard config: %v", err)
		return nil, err
	}
	return config, nil
}

func FindDashboardConfigByUserID(userID primitive.ObjectID) (*models.DashboardConfig, error) {
	var config models.DashboardConfig
	err := db.GetDB().Collection("dashboard_configs").FindOne(context.TODO(), bson.M{"user_id": userID}).Decode(&config)
	return &config, err
}

func UpdateDashboardConfig(configID primitive.ObjectID, updatedData bson.M) error {
	_, err := db.GetDB().Collection("dashboard_configs").UpdateOne(context.TODO(), bson.M{"_id": configID}, bson.M{"$set": updatedData})
	return err
}
