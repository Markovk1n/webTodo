package service

import (
	"github.com/markovk1n/webTodo/internal/repository"
	"github.com/markovk1n/webTodo/pkg/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
type TodoList interface {
	Create(userId int, list models.TodoList) (int, error)
	GetAllLists(userId int) ([]models.TodoList, error)
	GetListById(userId, listId int) (models.TodoList, error)
	DeleteListById(userId, listId int) error
	Update(userId, listId int, input models.UpdateListInput) error
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
		TodoList:      NewTodoService(repos.TodoList),
	}
}
