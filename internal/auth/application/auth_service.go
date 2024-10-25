package application

import (
	"context"
	"errors"
	"prontuario/internal/crypt"
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
	user, err := service.repo.FindUserByEmail(ctx, username)

	if err != nil {
		return nil, err
	}

	crypt := crypt.NewCrypt()

	comparasion := crypt.CompareHashAndPassword(user.Password, password)

	if comparasion == nil {
		return user, nil
	}

	return nil, errors.New("invalid password")
}
