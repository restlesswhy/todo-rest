package repository

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

func NewRepository() *Repository {
	return &Repository{}
}