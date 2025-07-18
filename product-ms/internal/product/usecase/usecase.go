package usecase

import (
	"context"
	"time"

	"product-ms/internal/product/domain"
)

type ProductUseCase struct {
	repo domain.ProductRepository
}

func NewProductUseCase(r domain.ProductRepository) *ProductUseCase {
	return &ProductUseCase{repo: r}
}

func (uc *ProductUseCase) CreateProduct(ctx context.Context, p *domain.Product) (*domain.Product, error) {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return uc.repo.Create(ctx, p)
}

func (uc *ProductUseCase) ListProducts(ctx context.Context) ([]*domain.Product, error) {
	return uc.repo.GetAll(ctx)
}

func (uc *ProductUseCase) GetProductByID(ctx context.Context, id string) (*domain.Product, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *ProductUseCase) UpdateProduct(ctx context.Context, id string, p *domain.Product) (*domain.Product, error) {
	p.UpdatedAt = time.Now()
	return uc.repo.Update(ctx, id, p)
}

func (uc *ProductUseCase) DeleteProduct(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}
