package routes

import (
	"ryanali12/web_service/app/repository"
	"ryanali12/web_service/controllers"

	"github.com/gin-gonic/gin"
)

func InitRootRoute(repos *repository.Repositories, router *gin.Engine) {
	var ctrler controllers.PublicController = controllers.PublicController{Repos: repos}
	router.GET("/", ctrler.ReturnIndexPage)
}
