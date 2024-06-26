package service

import (
	"github.com/markovk1n/webTodo/internal/models"
	"github.com/markovk1n/webTodo/internal/repository"
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
	Create(userId, listId int, item models.TodoItem) (int, error)
	GetAll(userId, listId int) ([]models.TodoItem, error)
	GetItem(userId, itemId int) (models.TodoItem, error)
	DeleteItem(userId, itemId int) error
	Update(userId, itemId int, input models.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
