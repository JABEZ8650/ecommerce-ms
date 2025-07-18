package domain

import "context"

type ProductRepository interface {
	Create(ctx context.Context, product *Product) (*Product, error)
	GetAll(ctx context.Context) ([]*Product, error)
	GetByID(ctx context.Context, id string) (*Product, error)
	Update(ctx context.Context, id string, product *Product) (*Product, error)
	Delete(ctx context.Context, id string) error
}
