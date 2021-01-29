package server

import (
	"7youo-wms/app/model"
	"7youo-wms/app/request"
	"7youo-wms/global"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	RoleNotFound = errors.New("角色不存在")
	ParentRoleNotFound  = errors.New("上级角色不存在")
	HaveChildRolesCannotDelete  = errors.New("此角色存在子角色，不允许删除")
	HaveUsersCannotDelete  = errors.New("此角色存在要不过户，不允许删除")
)
//创建角色
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
//删除角色
func DeleteRole(roleId int) (bool, error) {
	role := GetRoleById(roleId)
	//找不到记录
	if role == nil {
		return false, RoleNotFound
	}
	//存在子角色
	if !errors.Is(global.G_DB.Where("p_id = ?", roleId).First(&model.SysRole{}).Error, gorm.ErrRecordNotFound) {
		return false, HaveChildRolesCannotDelete
	}
	//存在用户
	if !errors.Is(global.G_DB.Where("role_id = ?", roleId).First(&model.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return false, HaveUsersCannotDelete
	}
	err := global.G_DB.Delete(role).Error
	//数据操作时出现错误
	if err != nil {
		global.G_LOG.Error("DeleteRole error", zap.Any("err", err.Error()))
		return false, err
	}
	return true, nil
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


