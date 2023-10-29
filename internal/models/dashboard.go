package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DashboardConfig struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID        primitive.ObjectID `bson:"user_id" json:"user_id"`
	Clock         ClockConfig        `bson:"clock" json:"clock"`
	Weather       WeatherConfig      `bson:"weather" json:"weather"`
	Traffic       TrafficConfig      `bson:"traffic" json:"traffic"`
	EventCalendar EventCalendar      `bson:"event_calendar" json:"event_calendar"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at"`
}

type ClockConfig struct {
	Timezone string `bson:"timezone" json:"timezone"`
	Is24Hour bool   `bson:"is_24_hour" json:"is_24_hour"`
	IsShown  bool   `bson:"is_shown" json:"is_shown"`
}

type WeatherConfig struct {
	WeatherLatitude  float64 `bson:"weather_lat" json:"weather_lat"`
	WeatherLongitude float64 `bson:"weather_long" json:"weather_long"`
	IsShown          bool    `bson:"is_shown" json:"is_shown"`
}

type TrafficConfig struct {
	LocationLatitude  float64 `bson:"location_lat" json:"location_lat"`
	LocationLongitude float64 `bson:"location_long" json:"location_long"`
	Label             string  `bson:"label" json:"label"`
	IsShown           bool    `bson:"is_shown" json:"is_shown"`
}

type EventCalendar struct {
	IsConnected bool `bson:"is_connected" json:"is_connected"`
	IsShown     bool `bson:"is_shown" json:"is_shown"`
}
