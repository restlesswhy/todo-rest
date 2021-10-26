package service

import "github.com/restlesswhy/todo-rest/repository"

type Authorization interface {
	
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
	return &Service{}
}