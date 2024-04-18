package service

import (
	"github.com/markovk1n/webTodo/internal/repository"
	"github.com/markovk1n/webTodo/pkg/models"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, item models.TodoItem) (int, error) {
	_, err := s.listRepo.GetListById(userId, listId)
	if err != nil {
		// list doesn't exists
		return 0, err
	}
	return s.repo.Create(listId, item)

}

func (s *TodoItemService) GetAll(userId, listId int) ([]models.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoItemService) GetItem(userId, itemId int) (models.TodoItem, error) {
	return s.repo.GetItem(userId, itemId)
}

func (s *TodoItemService) DeleteItem(userId, itemId int) error {
	return s.repo.DeleteItem(userId, itemId)
}

func (s *TodoItemService) Update(userId, itemId int, input models.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}
