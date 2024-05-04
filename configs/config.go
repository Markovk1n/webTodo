package configs

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/markovk1n/webTodo/internal/models"
)

func InitConfigs() (models.AppConfigs, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return models.AppConfigs{}, err
	}
	database := models.Postgres{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}

	token := os.Getenv("TOKEN")
	return models.AppConfigs{Postgres: database, Token: token}, nil
}
