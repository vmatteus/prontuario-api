package domain

import "context"

type UserRepositoryInterface interface {
	FindUserByEmail(ctx context.Context, email string) (*UserModel, error)
	CreateUser(ctx context.Context, user *UserModel) (*UserModel, error)
}
