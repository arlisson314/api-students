package main

import (
	"log"

	"github.com/arlisson314/api-students/api"
)

func main() {

	server := api.NewServer()
	server.ConfigureRoutes()
	server.StartServer()
	if err := server.StartServer(); err != nil {
		log.Fatal(err)
	}
}
