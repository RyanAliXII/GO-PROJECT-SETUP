package routes

import (
	"ryanali12/web_service/app/repository"
	"ryanali12/web_service/controllers"

	"github.com/gin-gonic/gin"
)

func InitLoginRoutes(repos *repository.Repositories, router *gin.Engine) {
	routerGroup := router.Group("/login")
	var ctrler controllers.LoginController = controllers.LoginController{Repos: repos}
	routerGroup.GET("/", ctrler.ReturnLoginPage)
}
