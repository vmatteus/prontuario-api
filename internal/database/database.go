package database

import (
	"prontuario/configs"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	config := configs.Get()
	switch config.Database.Driver {
	case "sqlite":
		return sqliteInit(config.Database.SqlLite)
	case "postgres":
		return postgresInit(config.Database.Postgres)
	default:
		return nil
	}
}

func sqliteInit(sqlLite *configs.SqlLite) error {
	db, err := gorm.Open(sqlite.Open(sqlLite.DbPath), &gorm.Config{})
	if err != nil {
		log.Error().Msg("failed to connect database: " + err.Error())
	}
	DB = db
	return nil
}

func postgresInit(postgres *configs.Postgres) error {
	return nil
}
