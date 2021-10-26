package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
	
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
	return &Repository{}
}