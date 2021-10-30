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

func (s *ListService) GetListById(userId int, idStr int) (todorest.List, error) {
	return s.repo.GetListById(userId, idStr)
}

func (s *ListService) DeleteList(userId int, idStr int) error {
	return s.repo.DeleteList(userId, idStr)
}

