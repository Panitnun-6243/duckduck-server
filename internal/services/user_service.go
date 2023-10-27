package services

import (
	"errors"
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"github.com/Panitnun-6243/duckduck-server/internal/repositories"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func RegisterUser(user *models.User) (*models.User, error) {
	// Check if device is available
	if !repositories.IsDeviceAvailable(user.DeviceCode) {
		return nil, errors.New("device code is not available or already used")
	}

	err := validate.Struct(user)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	return repositories.CreateUser(user)
}

func LoginUser(email, password string) (string, error) {
	// ... logic for user login ...
	// This will include verifying hashed passwords, generating JWT tokens, etc.
}
