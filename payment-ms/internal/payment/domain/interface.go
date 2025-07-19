package domain

import "context"

type PaymentRepository interface {
	CreatePayment(ctx context.Context, payment *Payment) (*Payment, error)
	GetPaymentByID(ctx context.Context, id string) (*Payment, error)
	GetAllPayments(ctx context.Context) ([]*Payment, error)
	UpdatePayment(ctx context.Context, id string, payment *Payment) (*Payment, error)
	DeletePayment(ctx context.Context, id string) error
}

type PaymentUseCase interface {
	CreatePayment(ctx context.Context, req *CreatePaymentRequest) (*Payment, error)
	GetPaymentByID(ctx context.Context, id string) (*Payment, error)
	GetAllPayments(ctx context.Context) ([]*Payment, error)
	UpdatePayment(ctx context.Context, id string, req *UpdatePaymentRequest) (*Payment, error)
	DeletePayment(ctx context.Context, id string) error
}
