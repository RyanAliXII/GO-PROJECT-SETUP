package routes

import (
	"ryanali12/web_service/controllers"
	"ryanali12/web_service/repository"

	"github.com/gin-gonic/gin"
)

func InitChatRoutes(r *gin.Engine, repos *repository.Repositories) {
	var controller controllers.ChatController = controllers.ChatController{}
	controller.UseRepos(repos)

	chatRoute := r.Group("/chat")
	chatRoute.GET("/", controller.Render)

}
