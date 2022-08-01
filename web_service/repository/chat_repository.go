package repository

import (
	"github.com/jmoiron/sqlx"
)

type chatRepository struct {
	db *sqlx.DB
}

type ChatRepository interface {
}
