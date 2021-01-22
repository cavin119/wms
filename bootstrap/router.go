package bootstrap

import (
	"7youo-wms/middlewares"
	"7youo-wms/router"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	var Router = gin.Default()
	PublicGroup := Router.Group("/api")
	{
		router.PublicRouter(PublicGroup)
	}
	PrivateGroup := Router.Group("/api")
	PrivateGroup.Use(middlewares.JWTAuth())
	{
		router.BaseRouter(PrivateGroup)
	}
	return Router
}
