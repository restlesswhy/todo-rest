package service

import (
	todorest "github.com/restlesswhy/todo-rest"
	"github.com/restlesswhy/todo-rest/repository"
)

type ItemService struct {
	itemRepo repository.Itemlist
	listRepo repository.Todolist
}

func NewItemService(itemRepo repository.Itemlist, listRepo repository.Todolist) *ItemService{
	return &ItemService{itemRepo: itemRepo, listRepo: listRepo}
}

func (s *ItemService) CreateItem(userId, listId int, itemInput todorest.Item) (int, error) {
	_, err := s.listRepo.GetListById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.itemRepo.CreateItem(listId, itemInput)
}

func (s *ItemService) GetAllItems(userId, listId int) ([]todorest.Item, error) {
	_, err := s.listRepo.GetListById(userId, listId)
	if err != nil {
		return nil, err
	}

	return s.itemRepo.GetAllItems(userId, listId)
}

func (s *ItemService) GetItemById(userId, itemId int) (todorest.Item, error) {
	return s.itemRepo.GetItemById(userId, itemId)
}

func (s *ItemService) DeleteItem(userId, itemId int) (error) {
	return s.itemRepo.DeleteItem(userId, itemId)
}