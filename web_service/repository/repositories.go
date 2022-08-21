package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	UserRepository UserRepository
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return userRepository{
		db: db,
	}
}
