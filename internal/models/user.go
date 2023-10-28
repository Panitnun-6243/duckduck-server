package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email      string             `bson:"email" json:"email" validate:"required,email"`
	Password   string             `bson:"password" json:"password" validate:"required"`
	Name       string             `bson:"name" json:"name" validate:"required"`
	DeviceCode string             `bson:"device_code" json:"device_code" validate:"required"`
	AvatarURL  string             `bson:"avatar_url" json:"avatar_url" validate:"required,url"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}

type UserProfile struct {
	Email     string `bson:"email" json:"email" validate:"required,email"`
	Name      string `bson:"name" json:"name" validate:"required"`
	AvatarURL string `bson:"avatar_url" json:"avatar_url" validate:"required,url"`
}
