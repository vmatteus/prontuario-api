package domain

import "context"

type UserRepositoryInterface interface {
	FindUserByUsernameAndPassword(ctx context.Context, username string, password string) (*UserModel, error)
	FindUserByEmail(ctx context.Context, email string) (*UserModel, error)
	CreateUser(ctx context.Context, user *UserModel) (*UserModel, error)
}
