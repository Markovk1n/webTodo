package service

import (
	"github.com/markovk1n/webTodo/internal/repository"
	"github.com/markovk1n/webTodo/pkg/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
}
type TodoList interface {
}
type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
