package routes

import (
	"ryanali12/web_service/app/repository"
	"ryanali12/web_service/controllers"

	"github.com/gin-gonic/gin"
)

func InitRootRoute(repos *repository.Repositories, router *gin.Engine) {
	var ctrler controllers.BaseController = controllers.BaseController{Repos: repos}
	router.GET("/", ctrler.ReturnIndexPage)
	router.GET("/product", ctrler.ReturnProductPage)
}
