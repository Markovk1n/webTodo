package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/markovk1n/webTodo/pkg/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
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

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: newAuthPostgres(db),
		TodoList:      NewTodoListRepository(db),
	}
}
