package resource

import (
	"7youo-wms/app/request"
	"7youo-wms/app/response"
	"7youo-wms/app/server"
	"7youo-wms/global"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//创建角色
func CreateRole(c *gin.Context) {
	var request request.CreateRole
	_ = c.ShouldBindJSON(&request)
	//数据校验
	errsMap, ok := response.RequestValidate(&request)
	if !ok {
		response.FailWithValidate(errsMap, c)
		return
	}
	//创建角色
	role, err := server.CreateRole(&request)
	if err != nil {
		//上级不存在
		if errors.Is(err, server.ParentRoleNotFound) {
			response.FailWithValidate(gin.H{
				"p_id": err.Error(),
			}, c)
			return
		}
		//插入数据失败
		global.G_LOG.Error("[CreateRole]创建角色失败，插入数据库失败：", zap.Any("err", err.Error()))
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OkWithDetailed(role, "创建角色成功", c)
}
