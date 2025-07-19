package mongo

import (
	"context"
	"errors"
	"payment-ms/internal/payment/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type paymentRepository struct {
	collection *mongo.Collection
}

func NewPaymentRepository(col *mongo.Collection) domain.PaymentRepository {
	return &paymentRepository{collection: col}
}

func (r *paymentRepository) CreatePayment(ctx context.Context, payment *domain.Payment) (*domain.Payment, error) {
	objectID := primitive.NewObjectID()
	payment.ID = objectID

	_, err := r.collection.InsertOne(ctx, payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (r *paymentRepository) GetPaymentByID(ctx context.Context, id string) (*domain.Payment, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid payment ID")
	}

	var payment domain.Payment
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&payment)
	if err != nil {
		return nil, errors.New("payment not found")
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
		return nil, errors.New("invalid payment ID")
	}

	update := bson.M{
		"$set": bson.M{
			"amount":    payment.Amount,
			"userId":    payment.UserID,
			"orderId":   payment.OrderID,
			"status":    payment.Status,
			"updatedAt": time.Now(),
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	result := r.collection.FindOneAndUpdate(ctx, bson.M{"_id": objID}, update, opts)
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, errors.New("payment not found")
		}
		return nil, result.Err()
	}

	var updated domain.Payment
	if err := result.Decode(&updated); err != nil {
		return nil, err
	}

	return &updated, nil
}

func (r *paymentRepository) DeletePayment(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid payment ID")
	}

	res, err := r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return errors.New("failed to delete payment")
	}
	if res.DeletedCount == 0 {
		return errors.New("payment not found")
	}

	return nil
}
