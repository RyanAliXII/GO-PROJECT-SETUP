package main

import (
	"net/http"
	"ryanali12/web_service/web"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("web/pages/**/*")
	r.Static("/static", "./web/static")
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome",
		})
		return
	})
	web.Init(r)
	r.Run()
}
