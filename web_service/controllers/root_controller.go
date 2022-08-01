package controllers

import (
	"net/http"
	"ryanali12/web_service/repository"

	"github.com/gin-gonic/gin"
)

type RootController struct {
	repos *repository.Repositories
}

func (ctrler *RootController) UseRepos(repos *repository.Repositories) {
	ctrler.repos = repos
}
func (ctrler *RootController) Render(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signin/index.html", nil)
}
