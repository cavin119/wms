package router

import (
	"7youo-wms/app/resource"
	"github.com/gin-gonic/gin"
)

func PublicRouter(router *gin.RouterGroup) (r gin.IRouter) {
	publicRouter := router.Group("/v1")
	{
		publicRouter.POST("/login", resource.Login)
		publicRouter.POST("/register", resource.Register)
	}
	return publicRouter
}
