package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	todorest "github.com/restlesswhy/todo-rest"
)

type ItemRepository struct {
	db *sqlx.DB
}

func NewItemRepository(db *sqlx.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) CreateItem(listId int, itemInput todorest.Item) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	itemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoItemTable)
	row := tx.QueryRow(itemQuery, itemInput.Title, itemInput.Description)

	if err := row.Scan(&itemId); err != nil {
		tx.Rollback()
		return 0, err
	}

	listItemQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listItemTable)
	_, err = tx.Exec(listItemQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *ItemRepository) GetAllItems(userId, listId int) ([]todorest.Item, error) {
	var items []todorest.Item

	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti 
							INNER JOIN %s li ON ti.id = li.item_id 
							INNER JOIN %s ul ON li.list_id = ul.list_id WHERE ul.user_id = $1 AND li.list_id = $2`, 
							todoItemTable, listItemTable, userListTable)

	if err := r.db.Select(&items, query, userId, listId); err != nil {
		return nil, err
	}

	return items, nil
}


func (r *ItemRepository) GetItemById(userId, itemId int) (todorest.Item, error) {
	var input todorest.Item

	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti 
							INNER JOIN %s li ON ti.id = li.item_id 
							INNER JOIN %s ul ON ul.list_id = li.list_id WHERE ul.user_id = $1 AND ti.id = $2`, todoItemTable, listItemTable, userListTable)
	err := r.db.Get(&input, query, userId, itemId)

	return input, err
}

func (r *ItemRepository) DeleteItem(userId, itemId int) (error) {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul 
							WHERE ti.id = li.item_id 
							AND ul.list_id = li.list_id 
							AND ul.user_id = $1 
							AND ti.id = $2`, todoItemTable, listItemTable, userListTable)
	_, err := r.db.Exec(query, userId, itemId)
	return err
}
