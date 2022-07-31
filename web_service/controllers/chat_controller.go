package controllers

import (
	"net/http"
	"ryanali12/web_service/repository"

	"github.com/gin-gonic/gin"
)

type ChatController struct {
	repos *repository.Repositories
}

func (ctrler *ChatController) UseRepos(repos *repository.Repositories) {
	ctrler.repos = repos
}
func (ctrler *ChatController) Render(ctx *gin.Context) {
	// user := models.User{
	// 	Firstname: "Ryan",
	// 	Lastname:  "Ali",
	// 	Email:     "ryanali456@gmail.com",
	// 	Password:  "test123",
	// }
	// ctrler.repos.UserRepository.CreateUser(user)
	ctx.HTML(http.StatusOK, "chat/index.html", nil)
}
