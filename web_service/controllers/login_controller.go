package controllers

import (
	"net/http"
	"ryanali12/web_service/app/repository"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	Repos *repository.Repositories
}

func (ctrl *LoginController) ReturnLoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "public/login.html", nil)
}
