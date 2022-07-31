package repository

import "database/sql"

type Repositories struct {
	ChatRepository ChatRepository
	UserRepository UserRepository
}

func NewUserRepository(db *sql.DB) UserRepository {
	return userRepository{
		db: db,
	}
}
func NewChatRepository(db *sql.DB) ChatRepository {
	return chatRepository{
		db: db,
	}
}
