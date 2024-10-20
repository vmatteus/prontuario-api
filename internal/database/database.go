package database

import (
	"fmt"
	"prontuario/configs"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	config := configs.Get()
	switch config.Database.Driver {
	case "sqlite":
		log.Info().Msg("connecting to sqlite database")
		return sqliteInit(config.Database.SqlLite)
	case "postgres":
		log.Info().Msg("connecting to postgres database")
		return postgresInit(config.Database.Postgres)
	default:
		log.Info().Msg("no database driver found: " + config.Database.Driver)
		return nil
	}
}

func sqliteInit(sqlLiteConfig *configs.SqlLite) error {
	db, err := gorm.Open(sqlite.Open(sqlLiteConfig.DbPath), &gorm.Config{})
	if err != nil {
		log.Error().Msg("failed to connect database: " + err.Error())
	}
	DB = db

	if err := Migrate(DB); err != nil {
		log.Error().Msg("failed to migrate: " + err.Error())
	}

	return nil
}

func postgresInit(postgresConfig *configs.Postgres) error {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		postgresConfig.Host, postgresConfig.User, postgresConfig.Password, postgresConfig.Database, postgresConfig.Port, postgresConfig.Sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Error().Msg("failed to connect database: " + err.Error())
	}
	DB = db

	if err := Migrate(DB); err != nil {
		log.Error().Msg("failed to migrate: " + err.Error())
	}

	return nil
}
