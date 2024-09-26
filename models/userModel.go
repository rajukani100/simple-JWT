package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	FirstName    string             `json:"first_name" bson:"first_name" validate:"required,min=3"`
	LastName     string             `json:"last_name" bson:"last_name" validate:"required,min=2"`
	Email        string             `json:"email" bson:"email" validate:"required,email"`
	HashPassword string             `json:"hash_password" bson:"hash_password" validate:"required,min=8"`
	PhoneNo      string             `json:"phone_no" bson:"phone_no" validate:"required,len=10,numeric"`
	Age          uint8              `json:"age" bson:"age" validate:"required,gte=18,lte=100"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
}
