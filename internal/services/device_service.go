package services

import (
	"errors"
	"github.com/Panitnun-6243/duckduck-server/internal/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeviceLogin(deviceCode, secret string) (string, error) {
	device, err := repositories.FindDeviceByCode(deviceCode)
	if err != nil {
		return "", err
	}
	if device == nil || device.Secret != secret {
		return "", errors.New("invalid device secret or device not exist")
	}

	// Assuming device is bound to a user, fetch the user here
	user, err := repositories.FindUserByDeviceCode(deviceCode)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("no user bound to this device")
	}

	token, err := generateToken(user, true)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetDeviceCodeByUserID returns the device code of the user
func GetDeviceCodeByUserID(userID primitive.ObjectID) (string, error) {
	deviceCode, err := repositories.FindDeviceCodeByUserID(userID)
	if err != nil {
		return "", err
	}
	return deviceCode, nil
}
