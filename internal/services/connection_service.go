package services

import (
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"github.com/Panitnun-6243/duckduck-server/internal/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetConnectionStatusByUser(userID primitive.ObjectID) (bool, error) {
	return repositories.FindConnectionStatusByUserID(userID)
}

func GetConnectionStatusByIDAndUserID(connID, userID primitive.ObjectID) (*models.LightControl, error) {
	return repositories.FindConnectionStatusByIDAndUserID(connID, userID)
}

func UpdateUserConnectionStatus(connID primitive.ObjectID, isConnected bool) error {
	return repositories.UpdateConnectionStatus(connID, isConnected)
}
