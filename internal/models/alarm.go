package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Alarm struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID                primitive.ObjectID `bson:"user_id" json:"user_id"`
	BedTime               *TimeModel         `bson:"bed_time" json:"bed_time"`
	WakeUpTime            *TimeModel         `bson:"wake_up_time" json:"wake_up_time"`
	Description           string             `bson:"description" json:"description"`
	IsActive              ActiveStatus       `bson:"is_active" json:"is_active"`
	RepeatDays            []string           `bson:"repeat_days" json:"repeat_days"`
	Sunrise               Sunrise            `bson:"sunrise" json:"sunrise"`
	CurrentAlarmSound     string             `bson:"current_alarm_sound" json:"current_alarm_sound"`
	CurrentAlarmSoundPath string             `bson:"current_alarm_sound_path" json:"current_alarm_sound_path"`
	Volume                float64            `bson:"volume" json:"volume"`
	SnoozeTime            int                `bson:"snooze_time" json:"snooze_time"`
	CreatedAt             time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt             time.Time          `bson:"updated_at" json:"updated_at"`
}

type TimeModel struct {
	Hours   int `bson:"hours" json:"hours" validate:"required,min=0,max=23"`
	Minutes int `bson:"minutes" json:"minutes" validate:"required,min=0,max=59"`
}

type ActiveStatus struct {
	Status     bool      `bson:"status" json:"status"`
	DateActive time.Time `bson:"date_active" json:"date_active"`
}

type Sunrise struct {
	StartTime *TimeModel `bson:"start_time" json:"start_time"`
	PeakTime  *TimeModel `bson:"peak_time" json:"peak_time"`
}

type TriggerAlarmRequest struct {
	ID string `json:"id"`
}

// CustomAlarmSound represents a user-uploaded custom alarm sound.
type CustomAlarmSound struct {
	UserID primitive.ObjectID `bson:"user_id" json:"user_id"`
	Sounds []SoundDetail      `bson:"sounds" json:"sounds"`
}

// PresetAlarmSound represents a preset alarm sound provided by the system.
type PresetAlarmSound struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name string             `bson:"name" json:"name"`
	Path string             `bson:"path" json:"path"`
}

// SoundDetail represents details of a sound.
type SoundDetail struct {
	Name string `bson:"name" json:"name"`
	Path string `bson:"path" json:"path"`
}
