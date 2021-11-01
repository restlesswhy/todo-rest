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
