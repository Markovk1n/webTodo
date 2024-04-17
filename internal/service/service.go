package service

import "github.com/markovk1n/webTodo/internal/repository"

type Authorization interface {
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
	return &Service{}
}
