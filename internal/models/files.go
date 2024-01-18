package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	ActiveStatus  = "active"
	DeletedStatus = "deleted"
)

type FileModel struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"file_name,omitempty" bson:"file_name"`
	CheckSum  string             `json:"check_sum,omitempty" bson:"check_sum"`
	Status    string             `json:"status,omitempty" bson:"status"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at"`
	DeletedAt time.Time          `json:"deleted_at,omitempty" bson:"deleted_at"`
}
