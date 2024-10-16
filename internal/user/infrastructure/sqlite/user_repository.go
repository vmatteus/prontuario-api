package sqlite

import (
	"context"
	"prontuario/internal/user/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindUserByUsernameAndPassword(ctx context.Context, username string, password string) (*domain.UserModel, error) {
	var user domain.UserModel
	result := r.db.Where("username = ? AND password = ?", username, password).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
