package infrastructure

import (
	"prontuario/configs"
	"prontuario/internal/database"
	"prontuario/internal/user/domain"
	infrastructureGorm "prontuario/internal/user/infrastructure/gorm"
)

func GetUserRepository() domain.UserRepositoryInterface {
	config := configs.Get()
	switch config.Database.Driver {
	case "sqlite":
		return infrastructureGorm.NewUserRepository(database.DB)
	case "postgres":
		return infrastructureGorm.NewUserRepository(database.DB)
	default:
		return nil
	}
}
