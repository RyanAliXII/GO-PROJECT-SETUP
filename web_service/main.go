package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"ryanali12/web_service/db"
	"ryanali12/web_service/repository"
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
	var repositories repository.Repositories = InitRepositories(connection)
	r.LoadHTMLFiles(getHTMLFiles()...)
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

func InitRepositories(db *sqlx.DB) repository.Repositories {
	return repository.Repositories{
		UserRepository: repository.NewUserRepository(db),
	}
}
func getHTMLFiles() []string {
	templateList := []string{}
	filepath.Walk("./templates", func(path string, info fs.FileInfo, err error) error {

		if !info.IsDir() {
			fileExtension := filepath.Ext(path)
			if fileExtension == ".html" {
				templateList = append(templateList, path)
			}

		}
		return nil
	})
	return templateList
}
