package service

import (
	todorest "github.com/restlesswhy/todo-rest"
	"github.com/restlesswhy/todo-rest/repository"
)

type Authorization interface {
	CreateUser(user todorest.User) (int, error)
}

type Todolist interface {
	
}

type Itemlist interface {
	
}

type Service struct {
	Authorization
	Todolist
	Itemlist
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthSerice(repo.Authorization),
	}
}