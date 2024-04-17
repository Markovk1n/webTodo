package main

import (
	"github.com/markovk1n/webTodo/internal/handler"
	"github.com/markovk1n/webTodo/internal/repository"
	"github.com/markovk1n/webTodo/internal/server"
	"github.com/markovk1n/webTodo/internal/service"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()

	repos := repository.NewRepository()
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(server.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("Server cant run ", err)
	}

}
