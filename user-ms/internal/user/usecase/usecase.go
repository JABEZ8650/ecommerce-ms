package usecase

import (
	"context"
	"time"
	"user-ms/internal/user/domain"
)

type userUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) domain.UserUseCase {
	return &userUseCase{repo: repo}
}

func (uc *userUseCase) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()
	return uc.repo.CreateUser(ctx, user)
}

func (uc *userUseCase) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	return uc.repo.GetUserByID(ctx, id)
}

func (uc *userUseCase) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	return uc.repo.GetAllUsers(ctx)
}

func (uc *userUseCase) UpdateUser(ctx context.Context, id string, user *domain.User) (*domain.User, error) {
	user.UpdatedAt = time.Now().Unix()
	return uc.repo.UpdateUser(ctx, id, user)
}

func (uc *userUseCase) DeleteUser(ctx context.Context, id string) error {
	return uc.repo.DeleteUser(ctx, id)
}
