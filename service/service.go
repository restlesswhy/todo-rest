package service

import (
	"github.com/restlesswhy/todo-rest"
	"github.com/restlesswhy/todo-rest/repository"
)

type Authorization interface {
	CreateUser(user todorest.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Todolist interface {
	CreateList(userId int, list todorest.List) (int, error)
	GetAllLists(UserId int) ([]todorest.List, error)
	GetListById(userId int, idStr int) (todorest.List, error)
	DeleteList(userId int, idStr int) error
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
		Todolist: NewListService(repo.Todolist),
	}
}