package main

import (
	"fmt"
	"os"

	"ryanali12/web_service/app/pkg/loadtmpl"
	"ryanali12/web_service/app/repository"
	"ryanali12/web_service/db"
	routes "ryanali12/web_service/routes/web"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	godotenv.Load(".env")
	var connection *sqlx.DB = InitDBConnection()
	db.InitDb(connection)
	var repositories repository.Repositories = repository.NewRepositories(connection)
	r.LoadHTMLFiles(loadtmpl.LoadHTMLFiles("./templates")...)
	r.Static("/public", "./public")
	routes.RegisterWeb(&repositories, r)
	r.Run()
}

func InitDBConnection() *sqlx.DB {

	DB_DRIVER := os.Getenv("DB_DRIVER")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")

	CONNECTION_STRING := fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT)
	fmt.Println(CONNECTION_STRING)
	connection, connectErr := sqlx.Open(DB_DRIVER, CONNECTION_STRING)
	if connectErr != nil {
		panic(connectErr.Error())
	}
	return connection
}
