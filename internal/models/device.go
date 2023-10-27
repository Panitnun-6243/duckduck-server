package models

type DeviceRegistrationCode struct {
	ID     string `bson:"_id" json:"id"`
	Code   string `bson:"code" json:"code"`
	UsedBy string `bson:"used_by" json:"used_by,omitempty"` // This will be a user ID.
}
