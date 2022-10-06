package routes

import (
	"ryanali12/web_service/app/repository"
	routes "ryanali12/web_service/routes/web/public"

	"github.com/gin-gonic/gin"
)

func RegisterWeb(r *repository.Repositories, router *gin.Engine) {
	routes.InitRootRoute(r, router)
	routes.InitLoginRoutes(r, router)
}
