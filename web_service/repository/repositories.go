package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	ChatRepository ChatRepository
	UserRepository UserRepository
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return userRepository{
		db: db,
	}
}
func NewChatRepository(db *sqlx.DB) ChatRepository {
	return chatRepository{
		db: db,
	}
}
