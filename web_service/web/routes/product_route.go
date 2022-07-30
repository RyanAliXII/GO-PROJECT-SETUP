package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitProductRoutes(r *gin.Engine) {
	productGroup := r.Group("/products")

	productGroup.GET("/create", func(ctx *gin.Context) {
		
		ctx.HTML(http.StatusOK, "pages/create_product.html", nil)
		// ctx.JSON(http.StatusOK, gin.H{"test": "test"})
		return
	})

}
