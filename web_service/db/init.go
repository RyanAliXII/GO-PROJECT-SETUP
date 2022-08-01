package db

import (
	"database/sql"
	"fmt"
	"os"
)

func InitDb(connection *sql.DB) {
	//ALWAYS ADD IF NOT EXISTS
	//CREATE YOUR OWN CUSTOM SCRIPT HERE TO CREATE DATABASE AND TABLE ON APPLICATION RUN
	createAppDb(connection)
}

func createAppDb(connection *sql.DB) {
	var DB_NAME string = os.Getenv("DB_NAME")
	connection.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", DB_NAME))
	connection.Exec(fmt.Sprintf("USE %s", DB_NAME))

	CREATE_USER_TABLE_QUERY := `
			CREATE TABLE IF NOT EXISTS users (
			id INT NOT NULL AUTO_INCREMENT,
			firstname VARCHAR(100) NOT NULL,
			lastname VARCHAR(100) NOT NULL,
			email VARCHAR(100) NOT NULL,
			password VARCHAR(100) NOT NULL,
			PRIMARY KEY (id));
	`
	connection.Exec(CREATE_USER_TABLE_QUERY)

}
