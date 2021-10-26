package service

import (
	todorest "github.com/restlesswhy/todo-rest"
	"github.com/restlesswhy/todo-rest/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthSerice(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todorest.User) (int, error) {
	return s.repo.CreateUser(user)
}