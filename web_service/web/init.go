package web

import (
	"net/http"
	"ryanali12/web_service/repository"
	"ryanali12/web_service/routes"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine, repos *repository.Repositories) {
	//INITIALIZE WEB ROUTES HERE
	routes.InitChatRoutes(r, repos)
	routes.InitRootRoute(r, repos)
	r.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "layouts/404.html", nil)
	})
}
