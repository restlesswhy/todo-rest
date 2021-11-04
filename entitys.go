package todorest

import "errors"

type List struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"decription"`
}

type UserList struct {
	Id     int 
	UserId int
	ListId int
}

type Item struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateListInput struct {
	Title *string `json:"title"`
	Description *string `json:"description"`
}

func (u *UpdateListInput) Validate() error {
	if u.Title == nil && u.Description == nil {
		return errors.New("have no values for update")
	}
	return nil
}

type UpdateItemInput struct {
	Title *string `json:"title"`
	Description *string `json:"description"`
	Done *bool `json:"done"` 
}

func (u *UpdateItemInput) Validate() error {
	if u.Title == nil && u.Description == nil && u.Done == nil {
		return errors.New("have no values for update")
	}
	return nil
}