package service

import (
	"github.com/restlesswhy/todo-rest"
	"github.com/restlesswhy/todo-rest/repository"
)

type ListService struct {
	repo repository.List
}

func NewListService(repo repository.List) *ListService {
	return &ListService{repo: repo}
}

func (s *ListService) CreateList(userId int, list todorest.List) (int, error) {
	return s.repo.CreateList(userId, list)
}

func (s *ListService) GetAllLists(userId int) ([]todorest.List, error) {
	return s.repo.GetAllLists(userId)
}

func (s *ListService) GetListById(userId int, listId int) (todorest.List, error) {
	return s.repo.GetListById(userId, listId)
}

func (s *ListService) DeleteList(userId int, listId int) error {
	return s.repo.DeleteList(userId, listId)
}

func (s *ListService) UpdateList(userId int, listId int, input todorest.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateList(userId, listId, input)
}
