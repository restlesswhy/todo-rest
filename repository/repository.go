package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/restlesswhy/todo-rest"
)

type Authorization interface {
	CreateUser(user todorest.User) (int, error)
	GetUser(username, password string) (int, error)
}

type Todolist interface {
	CreateList(userId int, list todorest.List) (int, error)
	GetAllLists(UserId int) ([]todorest.List, error)
	GetListById(userId int, idList int) (todorest.List, error)
	DeleteList(userId int, idList int) error
	UpdateList(userId int, idList int, input todorest.UpdateListInput) error
}

type Itemlist interface {
	
}

type Repository struct {
	Authorization
	Todolist
	Itemlist
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Todolist: NewListRepository(db),
	}
}