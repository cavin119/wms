package resource

import (
	apiRequest "7youo-wms/app/request"
	"7youo-wms/app/response"
	"7youo-wms/app/server"
	"7youo-wms/global"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func GetMenuById(c *gin.Context) {
	request := c.Param("id")

	menuId, err := strconv.Atoi(request)
	if err != nil {
		response.FailWithValidate(map[string]string{
			"id": "菜单ID格式错误",
		}, c)
		return
	}
	menu, notFound := server.GetMenuById(menuId)
	if notFound != nil {
		response.FailWithNotFound(notFound.Error(), c)
	}
	response.OkWithData(menu, c)
}

func DelMenuById(c *gin.Context) {
	request := c.Param("id")

	menuId, err := strconv.Atoi(request)
	if err != nil {
		response.FailWithValidate(map[string]string{
			"id": "菜单ID格式错误",
		}, c)
		return
	}
	ok, dErr := server.DeleteMenu(menuId)
	if !ok {
		if errors.Is(dErr, server.RoleNotFound) {
			//删除资源不存在
			response.FailWithNotFound(dErr.Error(), c)
		} else if errors.Is(dErr, server.HaveChildMenusCannotDelete) {
			//有子菜单
			response.FailWithBadRequest(dErr.Error(), c)
		} else {
			//数据删除失败未知错误
			response.FailWithMsg(dErr.Error(), c)
		}
		return
	}
	response.OkWithMsg("删除菜单成功", c)

}

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
	errsMap, ok := response.RequestValidate(&request)
	if !ok {
		response.FailWithValidate(errsMap, c)
		return
	}
	menu, err := server.CreateMenu(&request)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OkWithData(menu, c)
}
