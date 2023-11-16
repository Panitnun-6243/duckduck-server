package services

import (
	"errors"
	"fmt"
	"github.com/Panitnun-6243/duckduck-server/config"
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"github.com/Panitnun-6243/duckduck-server/internal/repositories"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateToken(user *models.User, isDevice bool) (string, error) {
	var expTime time.Time
	if isDevice {
		// For devices, use a longer expiration time.
		expTime = time.Now().Add(876000 * time.Hour)
	} else {
		// For regular users, use a shorter expiration time.
		expTime = time.Now().Add(730 * time.Hour)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   user.ID.Hex(),
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: expTime.Unix(),
	})

	cfg := config.LoadConfig()
	return token.SignedString([]byte(cfg.JWTSecret))
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

	_, err = CreateDefaultLightControl(user.ID)
	if err != nil {
		log.Println(fmt.Sprintf("Error while creating default light control: %v", err))
		return nil, err
	}
	_, err = CreateDefaultDashboardConfig(user.ID)
	if err != nil {
		log.Println(fmt.Sprintf("Error while creating default dashboard configuration: %v", err))
		return nil, err
	}
	_, err = CreateDefaultSleepClinic(user.ID)
	if err != nil {
		log.Println(fmt.Sprintf("Error while creating default sleep clinic: %v", err))
		return nil, err
	}

	// Bind the user ID to the device
	err = repositories.BindUserToDevice(user.ID, user.DeviceCode)
	if err != nil {
		log.Println(fmt.Sprintf("Error while binding user to device: %v", err))
		return nil, err
	}

	return user, nil
}

func LoginUser(email, password string) (string, error) {
	user, err := repositories.FindUserByEmail(email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user not found")
	}

	if !checkPasswordHash(password, user.Password) {
		return "", errors.New("invalid password")
	}

	token, err := generateToken(user, false)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetUserInfo(userID string) (*models.User, error) {
	return repositories.FindUserByID(userID)
}

func UpdateUserProfile(userID, name, avatarURL string) (*models.User, error) {
	user, err := repositories.UpdateUserDetails(userID, name, avatarURL)
	if err != nil {
		return nil, err
	}
	return user, nil
}
