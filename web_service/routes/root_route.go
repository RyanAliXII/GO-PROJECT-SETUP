package routes

import (
	"ryanali12/web_service/controllers"
	"ryanali12/web_service/repository"

	"github.com/gin-gonic/gin"
)

func InitRootRoute(r *gin.Engine, repos *repository.Repositories) {
	rootController := controllers.RootController{}
	rootController.UseRepos(repos)
	rootGroup := r.Group("/")
	rootGroup.GET("/signin", rootController.Render)
}
