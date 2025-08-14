package main

import (
	"github.com/arlisson314/api-students/api"
	"github.com/rs/zerolog/log"
)

func main() {

	server := api.NewServer()
	server.ConfigureRoutes()
	server.StartServer()

	if err := server.StartServer(); err != nil {
		log.Fatal().Err(err).Msgf("Failed to start server %s", err.Error())
	}

}
