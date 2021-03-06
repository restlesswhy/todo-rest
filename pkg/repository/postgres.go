package repository

import (
	"fmt"
	
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

const (
	userTable = "users"
	todoListTable = "lists"
	userListTable = "users_lists"
	todoItemTable = "items"
	listItemTable = "lists_items"
)

type Config struct {
	Host string
	Port string
	User string
	Password string
	DBName string
	SSLMode string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", 
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName,cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}


	return db, nil
}