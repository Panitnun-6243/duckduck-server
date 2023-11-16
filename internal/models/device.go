package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Device struct {
	ID     primitive.ObjectID `bson:"_id" json:"id"`
	Code   string             `bson:"code" json:"code"`
	IsUsed bool               `bson:"is_used" json:"is_used,omitempty"`
	Secret string             `bson:"secret" json:"secret"`
	UserID primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
}

type DeviceLoginRequest struct {
	DeviceCode string `json:"device_code"`
	Secret     string `json:"secret"`
}
