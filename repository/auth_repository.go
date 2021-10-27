package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/restlesswhy/todo-rest"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user todorest.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthRepository) GetUser(username, password string) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 and password_hash=$2", userTable)
	err := r.db.Get(&id, query, username, password)

	return id, err
}