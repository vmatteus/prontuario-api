package database

import (
	"prontuario/internal/user/domain"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.UserModel{},
	)
}
