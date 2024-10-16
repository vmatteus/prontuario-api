package main

import (
	"prontuario/configs"
	"prontuario/internal/database"
	"prontuario/internal/logger"
	"prontuario/internal/server"

	"github.com/rs/zerolog/log"
)

func main() {
	_, err := configs.Init()
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	logger.Init()

	if err := database.Init(); err != nil {
		log.Error().Msg(err.Error())
		return
	}

	server.Init()
}
