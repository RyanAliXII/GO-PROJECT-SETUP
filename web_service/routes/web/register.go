package routes

import (
	"net/http"
	"ryanali12/web_service/repository"

	"github.com/gin-gonic/gin"
)

func RegisterWeb(r *repository.Repositories, router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
}
