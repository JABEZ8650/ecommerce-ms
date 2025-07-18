package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" example:"64b22dd94c77c5b41f5a9b0d"`
	Name        string             `json:"name" bson:"name" example:"Water Bottle"`
	Description string             `json:"description" bson:"description" example:"Reusable plastic bottle"`
	Price       float64            `json:"price" bson:"price" example:"8.99"`
	CreatedAt   time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
