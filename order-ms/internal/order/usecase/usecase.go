package usecase

import (
	"context"
	"order-ms/internal/order/domain"
	"time"
)

// OrderUseCase is the interface that defines business logic for orders
type OrderUseCase interface {
	CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error)
	GetOrderByID(ctx context.Context, id string) (*domain.Order, error)
	GetOrders(ctx context.Context) ([]*domain.Order, error)
	UpdateOrder(ctx context.Context, id string, order *domain.Order) (*domain.Order, error)
	DeleteOrder(ctx context.Context, id string) error
}

// orderUseCase is the struct that implements OrderUseCase interface
type orderUseCase struct {
	repo domain.OrderRepository
}

// NewOrderUseCase creates a new instance of orderUseCase
func NewOrderUseCase(r domain.OrderRepository) OrderUseCase {
	return &orderUseCase{repo: r}
}

func (uc *orderUseCase) CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	order.CreatedAt = time.Now().Unix()
	order.UpdatedAt = time.Now().Unix()
	return uc.repo.Create(ctx, order)
}

func (uc *orderUseCase) GetOrderByID(ctx context.Context, id string) (*domain.Order, error) {
	return uc.repo.FindByID(ctx, id)
}

func (uc *orderUseCase) GetOrders(ctx context.Context) ([]*domain.Order, error) {
	return uc.repo.FindAll(ctx)
}

func (uc *orderUseCase) UpdateOrder(ctx context.Context, id string, order *domain.Order) (*domain.Order, error) {
	order.UpdatedAt = time.Now().Unix()
	return uc.repo.Update(ctx, id, order)
}

func (uc *orderUseCase) DeleteOrder(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}
