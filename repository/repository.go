package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/restlesswhy/todo-rest"
)

type Authorization interface {
	CreateUser(user todorest.User) (int, error)
	GetUser(username, password string) (int, error)
}

type List interface {
	CreateList(userId int, list todorest.List) (int, error)
	GetAllLists(UserId int) ([]todorest.List, error)
	GetListById(userId, listId int) (todorest.List, error)
	DeleteList(userId, listId int) error
	UpdateList(userId, listId int, input todorest.UpdateListInput) error
}

type Item interface {
	CreateItem(listId int, itemInput todorest.Item) (int, error)
	GetAllItems(userId, listId int) ([]todorest.Item, error)
	GetItemById(userId, itemId int) (todorest.Item, error)
	DeleteItem(userId, itemId int) (error)
	UpdateItem(userId, itemId int, input todorest.UpdateItemInput) error
}

type Repository struct {
	Authorization
	List
	Item
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		List: NewListRepository(db),
		Item: NewItemRepository(db),
	}
}