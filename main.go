package main

import (
	"github.com/markovk1n/webTodo/internal/handler"
	"github.com/markovk1n/webTodo/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	handler := new(handler.Handler)

	srv := new(server.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("Server cant run ", err)
	}

}
