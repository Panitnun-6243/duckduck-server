package services

import (
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"github.com/Panitnun-6243/duckduck-server/internal/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func CreateDefaultDashboardConfig(user primitive.ObjectID) (*models.DashboardConfig, error) {
	defaultDashboardConfig := &models.DashboardConfig{
		UserID: user,
		Clock: models.ClockConfig{
			Timezone: "Asia/Bangkok",
			Is24Hour: true,
			IsShown:  true,
		},
		Weather: models.WeatherConfig{
			WeatherLatitude:  13.736717,
			WeatherLongitude: 100.523186,
			IsShown:          false,
		},
		Traffic: models.TrafficConfig{
			LocationLatitude:  13.736717,
			LocationLongitude: 100.523186,
			Label:             "home",
			IsShown:           false,
		},
		EventCalendar: models.EventCalendar{
			IsConnected: false,
			IsShown:     false,
		},
		CreatedAt: time.Now(),
	}

	return repositories.CreateDashboardConfig(defaultDashboardConfig)
}

func GetDashboardConfigByUser(userID primitive.ObjectID) (*models.DashboardConfig, error) {
	return repositories.FindDashboardConfigByUserID(userID)
}

func UpdateUserDashboardConfig(configID primitive.ObjectID, updatedConfig *models.DashboardConfig) error {
	updatedConfig.UpdatedAt = time.Now()
	updateMap := bson.M{
		"clock":          updatedConfig.Clock,
		"weather":        updatedConfig.Weather,
		"traffic":        updatedConfig.Traffic,
		"event_calendar": updatedConfig.EventCalendar,
		"updated_at":     updatedConfig.UpdatedAt,
	}

	return repositories.UpdateDashboardConfig(configID, updateMap)
}
