package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SleepClinic struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID                 primitive.ObjectID `bson:"user_id" json:"user_id"`
	SleepStats             []SleepStat        `bson:"sleep_stats" json:"sleep_stats"`
	CurrentLullabySong     string             `bson:"current_lullaby_song" json:"current_lullaby_song"`
	CurrentLullabySongPath string             `bson:"current_lullaby_song_path" json:"current_lullaby_song_path"`
	DimLight               DimLight           `bson:"dim_light" json:"dim_light"`
}

type SleepStat struct {
	Date               string  `bson:"date" json:"date"`
	SleepDurationHours float64 `bson:"sleep_duration_hours" json:"sleep_duration_hours"`
}

type DimLight struct {
	IsActive bool `bson:"is_active" json:"is_active"`
	Duration int  `bson:"duration" json:"duration"` // Duration in minutes
}

type CustomLullabySong struct {
	UserID primitive.ObjectID `bson:"user_id" json:"user_id"`
	Songs  []LullabyDetail    `bson:"songs" json:"songs"`
}

type PresetLullabySong struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name" json:"name"`
	Path     string             `bson:"path" json:"path"`
	Category string             `bson:"category" json:"category"`
}

type LullabyDetail struct {
	Name     string `bson:"name" json:"name"`
	Path     string `bson:"path" json:"path"`
	Category string `bson:"category" json:"category"`
}
