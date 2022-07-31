package repository

import (
	"database/sql"
	"ryanali12/web_service/models"
)

type userRepository struct {
	db *sql.DB
}

func (repo userRepository) CreateUser(user models.User) {
	stmt, prepErr := repo.db.Prepare("INSERT INTO users(firstname,lastname,email,password)VALUES(?,?,?,?)")
	if prepErr != nil {
		panic(prepErr.Error())
	}
	stmt.Exec(user.Firstname, user.Lastname, user.Email, user.Password)
}

type UserRepository interface {
	CreateUser(models.User)
}
