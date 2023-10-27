package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DeviceRegistrationCode struct {
	ID     primitive.ObjectID `bson:"_id" json:"id"`
	Code   string             `bson:"code" json:"code"`
	IsUsed bool               `bson:"is_used" json:"is_used,omitempty"` // This will be a user ID.
}
