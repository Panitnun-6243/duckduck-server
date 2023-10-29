package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type LightControl struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID               primitive.ObjectID `bson:"user_id" json:"user_id"`
	Connected            bool               `bson:"connected" json:"connected"`
	SwitchStatus         string             `bson:"switch_status" json:"switch_status"`
	BrightnessPercentage float64            `bson:"brightness_percentage" json:"brightness_percentage"`
	ColorMode            string             `bson:"color_mode" json:"color_mode"`
	CctTemp              int                `bson:"cct_temp" json:"cct_temp"`
	RgbColor             string             `bson:"rgb_color" json:"rgb_color"`
	CreatedAt            time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt            time.Time          `bson:"updated_at" json:"updated_at"`
}
