package service

import (
	"github.com/restlesswhy/todo-rest"
	"github.com/restlesswhy/todo-rest/repository"
)

type ListService struct {
	repo repository.Todolist
}

func NewListService(repo repository.Todolist) *ListService {
	return &ListService{repo: repo}
}

func (s *ListService) CreateList(userId int, list todorest.List) (int, error) {
	return s.repo.CreateList(userId, list)
}

func (s *ListService) GetAllLists(userId int) ([]todorest.List, error) {
	return s.repo.GetAllLists(userId)
}

func (s *ListService) GetListById(userId int, idList int) (todorest.List, error) {
	return s.repo.GetListById(userId, idList)
}

func (s *ListService) DeleteList(userId int, idList int) error {
	return s.repo.DeleteList(userId, idList)
}

func (s *ListService) UpdateList(userId int, idList int, input todorest.UpdateListInput) error {
	return s.repo.UpdateList(userId, idList, input)
}
