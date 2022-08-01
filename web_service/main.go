package main

import (
	"database/sql"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"ryanali12/web_service/db"
	"ryanali12/web_service/repository"
	"ryanali12/web_service/web"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	godotenv.Load(".env")

	var connection *sql.DB = InitDBConnection()
	db.InitDb(connection)
	var repositories repository.Repositories = InitRepositories(connection)
	r.LoadHTMLFiles(getHTMLFiles()...)
	r.Static("/static", "./web/static")
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome",
		})
	})
	web.Init(r, &repositories)
	r.Run()
}

func InitDBConnection() *sql.DB {

	DB_DRIVER := os.Getenv("DB_DRIVER")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")

	CONNECTION_STRING := fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT)
	fmt.Println(CONNECTION_STRING)
	connection, connectErr := sql.Open(DB_DRIVER, CONNECTION_STRING)
	if connectErr != nil {
		panic(connectErr.Error())
	}
	return connection
}

func InitRepositories(db *sql.DB) repository.Repositories {

	return repository.Repositories{
		ChatRepository: repository.NewChatRepository(db),
		UserRepository: repository.NewUserRepository(db),
	}
}
func getHTMLFiles() []string {
	templateList := []string{}
	filepath.Walk("./web/templates", func(path string, info fs.FileInfo, err error) error {

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
