package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/restlesswhy/todo-rest"
	"github.com/sirupsen/logrus"
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

func (r *ListRepository) GetAllLists(userId int) ([]todorest.List, error) {
	var lists []todorest.List

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id=ul.list_id WHERE ul.user_id=$1", 
		todoListTable, userListTable)
	if err := r.db.Select(&lists, query, userId); err != nil {
		return nil, err
	}

	return lists, nil
}

func (r *ListRepository) GetListById(userId int, idList int) (todorest.List, error) {
	var list todorest.List

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id=ul.list_id WHERE ul.list_id=$1 and ul.user_id=$2", 
		todoListTable, userListTable)
	err := r.db.Get(&list, query, idList, userId)

	return list, err
}

func (r *ListRepository) DeleteList(userId int, idList int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE ul.list_id=tl.id and ul.user_id=$1 and ul.list_id=$2", 
		todoListTable, userListTable)
	_, err := r.db.Exec(query, userId, idList)
	return err
}

func (r *ListRepository) UpdateList(userId int, idList int, input todorest.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argsId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argsId))
		args = append(args, *input.Title)
		argsId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argsId))
		args = append(args, *input.Description)
		argsId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.user_id = $%d AND ul.list_id = $%d", 
		todoListTable, setQuery, userListTable, argsId, argsId+1)
	args = append(args, userId, idList)

	logrus.Debugf("updateQuery = %s", query)
	logrus.Debugf("args = %s", args)
	
	_, err := r.db.Exec(query, args...)

	return err
}

