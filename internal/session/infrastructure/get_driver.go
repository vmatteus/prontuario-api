package infrastructure

import (
	"prontuario/configs"
	"prontuario/internal/session/domain"
	infrastructureDatabase "prontuario/internal/session/infrastructure/database"
	infrastructureRedis "prontuario/internal/session/infrastructure/redis"
)

func GetSessionDriver() domain.SessionDriverInterface {
	config := configs.Get()
	switch config.Session.Driver {
	case "database":
		return infrastructureDatabase.NewSessionDriver()
	case "redis":
		return infrastructureRedis.NewSessionDriver()
	default:
		return nil
	}
}
