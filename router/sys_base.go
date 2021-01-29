package router

import (
	"7youo-wms/app/resource"
	"github.com/gin-gonic/gin"
)

func BaseRouter(router *gin.RouterGroup) (r gin.IRouter) {
	BaseRouter := router.Group("/v1")
	{
		BaseRouter.GET("/user/info", resource.UserInfo)
		//角色
		BaseRouter.POST("/role/create", resource.CreateRole)

		//角色
		BaseRouter.DELETE("/role/:id", resource.DeleteRole)
	}
	return BaseRouter
}
