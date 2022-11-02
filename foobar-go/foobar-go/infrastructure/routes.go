package infrastructure

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GinRouter struct {
	Gin *gin.Engine
}

func NewGinRouter() GinRouter {
	httpRouter := gin.Default()
	httpRouter.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"status": "up"})
	})

	return GinRouter{
		Gin: httpRouter,
	}
}
