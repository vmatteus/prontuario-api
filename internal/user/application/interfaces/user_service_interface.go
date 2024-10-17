package interfaces

import (
	"context"
	"prontuario/internal/user/domain"
)

type UserServiceInterface interface {
	Create(ctx context.Context, user *domain.UserModel) (*domain.UserModel, error)
}
