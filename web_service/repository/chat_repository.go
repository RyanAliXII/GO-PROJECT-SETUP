package repository

import "database/sql"

type chatRepository struct {
	db *sql.DB
}

type ChatRepository interface {
}
