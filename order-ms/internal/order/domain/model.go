package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CustomerID string             `bson:"customer_id" json:"customer_id" validate:"required"`
	ProductID  string             `bson:"product_id" json:"product_id" validate:"required"`
	Quantity   int                `bson:"quantity" json:"quantity" validate:"required,min=1"`
	Status     string             `bson:"status" json:"status" validate:"required,oneof=pending confirmed shipped delivered"`
	CreatedAt  int64              `bson:"created_at" json:"created_at"`
	UpdatedAt  int64              `bson:"updated_at" json:"updated_at"`
}
type OrderRepository interface {
	Create(ctx context.Context, order *Order) (*Order, error)
	FindByID(ctx context.Context, id string) (*Order, error)
	FindAll(ctx context.Context) ([]*Order, error)
	Update(ctx context.Context, id string, order *Order) (*Order, error)
	Delete(ctx context.Context, id string) error
}

type OrderUseCase interface {
	CreateOrder(ctx context.Context, order *Order) (*Order, error)
	GetOrderByID(ctx context.Context, id string) (*Order, error)
	GetOrders(ctx context.Context) ([]*Order, error)
	UpdateOrder(ctx context.Context, id string, order *Order) (*Order, error)
	DeleteOrder(ctx context.Context, id string) error
}
