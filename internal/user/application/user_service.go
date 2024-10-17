package application

import (
	"context"
	"prontuario/internal/user/domain"
	"prontuario/internal/user/infrastructure"
)

type UserService struct {
	repo domain.UserRepositoryInterface
}

func NewUserService() *UserService {
	return &UserService{repo: infrastructure.GetUserRepository()}
}

func (service *UserService) Create(ctx context.Context, user *domain.UserModel) (*domain.UserModel, error) {
	return service.repo.CreateUser(ctx, user)
}
