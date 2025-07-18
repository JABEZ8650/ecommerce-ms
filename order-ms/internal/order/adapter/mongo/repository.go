package mongo

import (
	"context"
	"errors"
	"time"

	"order-ms/internal/order/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// orderRepository is the struct that implements domain.OrderRepository
type orderRepository struct {
	collection *mongo.Collection
}

// NewOrderRepository creates a new instance of orderRepository
func NewOrderRepository(col *mongo.Collection) domain.OrderRepository {
	return &orderRepository{
		collection: col,
	}
}

func (r *orderRepository) Create(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	order.ID = primitive.NewObjectID()
	order.CreatedAt = time.Now().Unix()
	order.UpdatedAt = time.Now().Unix()

	_, err := r.collection.InsertOne(ctx, order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *orderRepository) FindByID(ctx context.Context, id string) (*domain.Order, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var order domain.Order
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&order)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &order, nil
}

func (r *orderRepository) FindAll(ctx context.Context) ([]*domain.Order, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []*domain.Order
	for cursor.Next(ctx) {
		var order domain.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *orderRepository) Update(ctx context.Context, id string, order *domain.Order) (*domain.Order, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	order.UpdatedAt = time.Now().Unix()

	update := bson.M{
		"$set": order,
	}

	_, err = r.collection.UpdateByID(ctx, objectID, update)
	if err != nil {
		return nil, err
	}

	return r.FindByID(ctx, id)
}

func (r *orderRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}
