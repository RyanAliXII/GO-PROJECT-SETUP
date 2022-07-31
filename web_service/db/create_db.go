package db

import (
	"database/sql"
	"fmt"
)

func CreateDBandTables(connection *sql.DB, name string) {
	_, CreateDBErr := connection.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", name))
	if CreateDBErr != nil {
		panic(CreateDBErr.Error())
	}

	//CREATE YOUR OWN CUSTOM SCRIPT HERE TO CREATE TABLE ON APPLICATION RUN
	//ALWAYS ADD IF NOT EXISTS

	connection.Exec(fmt.Sprintf("USE %s", name))

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
