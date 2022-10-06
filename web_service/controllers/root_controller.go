package controllers

import (
	"net/http"
	"ryanali12/web_service/app/repository"

	"github.com/gin-gonic/gin"
)

type PublicController struct {
	Repos *repository.Repositories
}

func (ctrl *PublicController) ReturnIndexPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "public/login.html", nil)
}

//
