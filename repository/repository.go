package repository

import (
	"github.com/jmoiron/sqlx"
	todorest "github.com/restlesswhy/todo-rest"
)

type Authorization interface {
	CreateUser(user todorest.User) (int, error)
}

type Todolist interface {
	
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
	}
}