package infrastructure

import (
	"prontuario/configs"
	"prontuario/internal/database"
	"prontuario/internal/user/domain"
	infrastructure "prontuario/internal/user/infrastructure/sqlite"
)

func GetUserRepository() domain.UserRepositoryInterface {
	config := configs.Get()
	switch config.Database.Driver {
	case "sqlite":
		return infrastructure.NewUserRepository(database.DB)
	default:
		return nil
	}
}
