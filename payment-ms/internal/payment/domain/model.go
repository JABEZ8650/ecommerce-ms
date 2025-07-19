package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    string             `json:"userId" bson:"userId"`
	OrderID   string             `json:"orderId" bson:"orderId"`
	Amount    float64            `json:"amount" bson:"amount"`
	Status    string             `json:"status" bson:"status"` // e.g., "pending", "completed", "failed"
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}
