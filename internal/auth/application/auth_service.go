package application

import (
	"context"
	"prontuario/internal/user/domain"
	"prontuario/internal/user/infrastructure"
)

type AuthService struct {
	repo domain.UserRepositoryInterface
}

func NewAuthService() *AuthService {
	return &AuthService{repo: infrastructure.GetUserRepository()}
}

func (service *AuthService) Login(ctx context.Context, username, password string) (*domain.UserModel, error) {
	return service.repo.FindUserByUsernameAndPassword(ctx, username, password)
}
