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

type List interface {
	CreateList(userId int, list todorest.List) (int, error)
	GetAllLists(UserId int) ([]todorest.List, error)
	GetListById(userId, listId int) (todorest.List, error)
	DeleteList(userId, listId int) error
	UpdateList(userId, listId int, input todorest.UpdateListInput) error
}

type Item interface {
	CreateItem(userId, listId int, itemInput todorest.Item) (int, error)
	GetAllItems(userId, listId int) ([]todorest.Item, error)
	GetItemById(userId, itemId int) (todorest.Item, error)
	DeleteItem(userId, itemId int) (error)
	UpdateItem(userId, itemId int, input todorest.UpdateItemInput) error
}

type Service struct {
	Authorization
	List
	Item
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthSerice(repo.Authorization),
		List: NewListService(repo.List),
		Item: NewItemService(repo.Item, repo.List),
	}
}