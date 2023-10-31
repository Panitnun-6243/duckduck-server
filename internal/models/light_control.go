package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type LightControl struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID               primitive.ObjectID `bson:"user_id" json:"user_id"`
	Connected            bool               `bson:"connected" json:"connected"`
	SwitchStatus         bool               `bson:"on" json:"on"`
	BrightnessPercentage float64            `bson:"brightness" json:"brightness"`
	ColorMode            string             `bson:"color_mode" json:"color_mode"`
	CctTemp              int                `bson:"temp" json:"temp"`
	HslColor             Hsl                `bson:"hsl_color" json:"hsl_color"`
	CreatedAt            time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt            time.Time          `bson:"updated_at" json:"updated_at"`
}

type Hsl struct {
	Hue        int `bson:"h" json:"h"`
	Saturation int `bson:"s" json:"s"`
	Lightness  int `bson:"l" json:"l"`
}
