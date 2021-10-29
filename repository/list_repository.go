package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/restlesswhy/todo-rest"
)

type ListRepository struct {
	db *sqlx.DB
}

func NewListRepository(db *sqlx.DB) *ListRepository {
	return &ListRepository{db: db}
}

func (r *ListRepository) CreateList(userId int, list todorest.List) (int, error) {
	ts, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var listId int
	listQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListTable)
	row := ts.QueryRow(listQuery, list.Title, list.Description)

	if err := row.Scan(&listId); err != nil {
		ts.Rollback()
		return 0, err
	}

	userListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", userListTable)
	_, err = ts.Exec(userListQuery, userId, listId)

	return listId, ts.Commit()
}