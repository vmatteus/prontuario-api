package interfaces

import (
	"context"
	"prontuario/internal/user/domain"
)

type AuthServiceInterface interface {
	Login(ctx context.Context, username, password string) (*domain.UserModel, string, error)
	HashPassword(password string) (string, error)
	CompareHashAndPassword(hashedPassword, password string) error
	ValidateToken(token string) (*domain.UserModel, error)
}
