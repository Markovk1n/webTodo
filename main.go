package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/markovk1n/webTodo/configs"
	"github.com/markovk1n/webTodo/internal/handler"
	"github.com/markovk1n/webTodo/internal/repository"
	"github.com/markovk1n/webTodo/internal/server"
	"github.com/markovk1n/webTodo/internal/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	log := logrus.New()

	if err := configs.InitConfig(); err != nil {
		log.Fatalf("cant read config %s", err)
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			log.Fatalf("Server cant run %s", err.Error())
		}
	}()
	log.Print("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Print("TodoApp Shutting Down")

	if err := srv.ShutDown(context.Background()); err != nil {
		log.Errorf("error occured on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close:", err.Error())
	}
}
