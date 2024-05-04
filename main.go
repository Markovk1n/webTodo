package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
	"github.com/markovk1n/webTodo/configs"
	"github.com/markovk1n/webTodo/internal/handler"
	"github.com/markovk1n/webTodo/internal/models"
	"github.com/markovk1n/webTodo/internal/repository"
	"github.com/markovk1n/webTodo/internal/server"
	"github.com/markovk1n/webTodo/internal/service"
	"github.com/sirupsen/logrus"
)

// @title Todo App API
// @version 1.0
// description API Server for TodoList Application

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	log := logrus.New()

	cfg, err := configs.InitConfigs()
	if err != nil {
		log.Fatalf("error with loading conf datas: %s", err.Error())

	}
	log.Println("Wait db creating")
	time.Sleep(10 * time.Second)

	db, err := repository.NewPostgresDB(models.Postgres{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		Username: cfg.Postgres.Username,
		Password: cfg.Postgres.Password,
		DBName:   cfg.Postgres.DBName,
		SSLMode:  cfg.Postgres.SSLMode,
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(server.Server)
	go func() {
		if err = srv.Run("8080", handler.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
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
