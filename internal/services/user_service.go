package services

import (
	"errors"
	"fmt"
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"github.com/Panitnun-6243/duckduck-server/internal/repositories"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// for login
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RegisterUser(user *models.User) (*models.User, error) {
	existingUser, err := repositories.FindUserByEmail(user.Email)
	if err != nil {
		log.Println(fmt.Sprintf("Error while finding user by email: %v", err))
		return nil, err
	}
	if existingUser != nil {
		log.Println("Email already registered")
		return nil, errors.New("email already registered")
	}

	if !repositories.IsDeviceAvailable(user.DeviceCode) {
		log.Println("Device code is not available or already used")
		return nil, errors.New("device code is not available or already used")
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	user, err = repositories.CreateUser(user)
	if err != nil {
		log.Println(fmt.Sprintf("Error while creating user: %v", err))
		return nil, err
	}

	err = repositories.MarkDeviceAsUsed(user.DeviceCode)
	if err != nil {
		log.Println(fmt.Sprintf("Error while marking device as used: %v", err))
		return nil, err
	}

	return user, nil
}
