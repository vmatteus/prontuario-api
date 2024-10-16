package interfaces

import (
	"context"
	"prontuario/internal/user/domain"
)

type AuthServiceInterface interface {
	Login(ctx context.Context, username, password string) (*domain.UserModel, error)
}
