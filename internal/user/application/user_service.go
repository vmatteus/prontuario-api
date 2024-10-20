package application

import (
	"context"
	"errors"
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

	_, err := service.repo.FindUserByEmail(ctx, user.Email)

	if err == nil {
		return user, errors.New("user already exists")
	}

	return service.repo.CreateUser(ctx, user)
}
