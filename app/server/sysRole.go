package server

import (
	"7youo-wms/app/model"
	"7youo-wms/app/request"
	"7youo-wms/global"
	"errors"
)

var (
	ParentRoleNotFound  = errors.New("上级角色不存在")
)

func CreateRole(request *request.CreateRole) (role *model.SysRole, err error){
	var newRole model.SysRole
	newRole.RoleName = request.RoleName
	if request.PId > 0 {
		pRole := GetRoleById(request.PId)
		//没获取上级role
		if pRole == nil {
			return nil, ParentRoleNotFound
		}
		newRole.PId = request.PId
	}

	err = global.G_DB.Create(&newRole).Error
	return &newRole, err
}

//通过id获取角色详情
//@param roleId int
//@return nil
//@return SysRole
func GetRoleById(roleId int) (role *model.SysRole)  {
	var model model.SysRole
	rows := global.G_DB.Where("id = ?", roleId).Find(&model).RowsAffected
	if rows == 0 {
		return nil
	}
	return &model
}


