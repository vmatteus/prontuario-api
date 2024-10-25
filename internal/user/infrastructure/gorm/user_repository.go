package infrastructure

import (
	"context"
	"errors"
	"prontuario/internal/user/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.UserModel) (*domain.UserModel, error) {

	result := r.db.Create(user)
	if result.Error != nil {
		return nil, errors.New("error creating user")
	}

	return user, nil
}

func (r *UserRepository) FindUserByEmail(ctx context.Context, email string) (*domain.UserModel, error) {
	var user domain.UserModel
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
