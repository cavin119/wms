package model

type SysRole struct {
	WMSModel
	PId int `json:"p_id" gorm:"comment:上级ID;default:0"`
	RoleName string `json:"role_name" gorm:"comment:角色名称"`
}