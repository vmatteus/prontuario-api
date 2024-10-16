package infrastructure

import (
	"prontuario/configs"
	"prontuario/internal/database"
	"prontuario/internal/user/domain"
	"prontuario/internal/user/infrastructure/sqlite"
)

func GetUserRepository() domain.UserRepositoryInterface {
	config := configs.Get()
	switch config.Database.Driver {
	case "sqlite":
		return sqlite.NewUserRepository(database.DB)
	default:
		return nil
	}
}
