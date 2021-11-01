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