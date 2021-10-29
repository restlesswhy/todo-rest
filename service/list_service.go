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