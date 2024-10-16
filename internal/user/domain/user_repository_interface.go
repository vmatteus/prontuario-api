package domain

import "context"

type UserRepositoryInterface interface {
	FindUserByUsernameAndPassword(ctx context.Context, username string, password string) (*UserModel, error)
}
