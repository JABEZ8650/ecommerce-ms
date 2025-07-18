package mongo

import (
	"context"
	"errors"
	"time"

	"payment-ms/internal/payment/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type paymentRepository struct {
	collection *mongo.Collection
}

func NewPaymentRepository(col *mongo.Collection) domain.PaymentRepository {
	return &paymentRepository{collection: col}
}

func (r *paymentRepository) CreatePayment(ctx context.Context, payment *domain.Payment) (*domain.Payment, error) {
	objectID := primitive.NewObjectID()
	payment.ID = objectID.Hex()
	payment.CreatedAt = time.Now().Unix()
	payment.UpdatedAt = time.Now().Unix()

	_, err := r.collection.InsertOne(ctx, payment)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *paymentRepository) GetPaymentByID(ctx context.Context, id string) (*domain.Payment, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var payment domain.Payment
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&payment)
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepository) GetAllPayments(ctx context.Context) ([]*domain.Payment, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var payments []*domain.Payment
	for cursor.Next(ctx) {
		var payment domain.Payment
		if err := cursor.Decode(&payment); err != nil {
			return nil, err
		}
		payments = append(payments, &payment)
	}
	return payments, nil
}

func (r *paymentRepository) UpdatePayment(ctx context.Context, id string, payment *domain.Payment) (*domain.Payment, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	payment.UpdatedAt = time.Now().Unix()
	update := bson.M{"$set": payment}

	_, err = r.collection.UpdateByID(ctx, objID, update)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *paymentRepository) DeletePayment(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	res, err := r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("payment not found")
	}
	return nil
}
