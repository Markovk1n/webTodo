package service

import (
	"github.com/markovk1n/webTodo/internal/repository"
	"github.com/markovk1n/webTodo/pkg/models"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list models.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAllLists(userId int) ([]models.TodoList, error) {
	return s.repo.GetAllLists(userId)
}

func (s *TodoListService) GetListById(userId, listId int) (models.TodoList, error) {
	return s.repo.GetListById(userId, listId)
}

func (s *TodoListService) DeleteListById(userId, listId int) error {
	return s.repo.DeleteListById(userId, listId)
}

func (s *TodoListService) Update(userId, listId int, input models.UpdateListInput) error {
	return s.repo.Update(userId, listId, input)
}
