package routes

import (
	"net/http"
	"ryanali12/web_service/app/repository"

	"github.com/gin-gonic/gin"
)

func RegisterWeb(r *repository.Repositories, router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/about", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "about.html", nil)
	})
}
