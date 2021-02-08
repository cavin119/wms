package resource

import (
	apiRequest "7youo-wms/app/request"
	"7youo-wms/app/response"
	"7youo-wms/app/server"
	"7youo-wms/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetMenus(c *gin.Context) {
	err, menus := server.GetMenuTree()
	if err!= nil {
		global.G_LOG.Error("GetMenus error:", zap.Any("err", err.Error()))
		response.FailWithMsg("获取菜单失败", c)
		return
	}
	response.OkWithData(menus, c)
}

func CreateMenu(c *gin.Context) {
	var request apiRequest.CreateMenu
	c.ShouldBindJSON(&request)
	fmt.Printf("CreateMenu request: %v", &request)
	errsMap, ok := response.RequestValidate(&request)
	if !ok {
		response.FailWithValidate(errsMap, c)
		return
	}
	response.Ok(c)
}
