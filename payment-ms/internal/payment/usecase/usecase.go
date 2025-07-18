package usecase

import (
	"context"
	"payment-ms/internal/payment/domain"
)

type paymentUseCase struct {
	repo domain.PaymentRepository
}

func NewPaymentUseCase(repo domain.PaymentRepository) domain.PaymentUseCase {
	return &paymentUseCase{repo: repo}
}

func (uc *paymentUseCase) CreatePayment(ctx context.Context, req domain.CreatePaymentRequest) (*domain.Payment, error) {
	payment := &domain.Payment{
		OrderID: req.OrderID,
		Amount:  req.Amount,
		Status:  req.Status,
	}
	return uc.repo.CreatePayment(ctx, payment)
}

func (uc *paymentUseCase) GetPaymentByID(ctx context.Context, id string) (*domain.Payment, error) {
	return uc.repo.GetPaymentByID(ctx, id)
}

func (uc *paymentUseCase) GetAllPayments(ctx context.Context) ([]*domain.Payment, error) {
	return uc.repo.GetAllPayments(ctx)
}

func (uc *paymentUseCase) UpdatePayment(ctx context.Context, id string, req domain.UpdatePaymentRequest) (*domain.Payment, error) {
	payment := &domain.Payment{
		Amount: req.Amount,
		Status: req.Status,
	}
	return uc.repo.UpdatePayment(ctx, id, payment)
}

func (uc *paymentUseCase) DeletePayment(ctx context.Context, id string) error {
	return uc.repo.DeletePayment(ctx, id)
}
