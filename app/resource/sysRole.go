package resource

import (
	"7youo-wms/app/request"
	"7youo-wms/app/response"
	"7youo-wms/app/server"
	"7youo-wms/global"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
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

func DeleteRole(c *gin.Context) {
	request := c.Param("id")

	roleId, err := strconv.Atoi(request)
	if err != nil {
		response.FailWithValidate(map[string]string{
			"id": "角色ID格式错误",
		}, c)
		return
	}
	ok, dErr := server.DeleteRole(roleId)
	if !ok {
		if errors.Is(dErr, server.RoleNotFound) {
			//删除资源不存在
			response.FailWithNotFound(dErr.Error(), c)
		} else if errors.Is(dErr, server.HaveChildRolesCannotDelete) {
			//有子角色
			response.FailWithBadRequest(dErr.Error(), c)
		} else {
			//数据删除失败未知错误
			response.FailWithMsg(dErr.Error(), c)
		}
		return
	}
	response.OkWithMsg("角色删除成功", c)
}
