package model

type SysUser struct {
	WMSModel
	Username string `json:"username" gorm:"comment:用户名"`
	Password string `json:"password" gorm:"comment:密码"`
	RoleId   int    `json:"role_id" gorm:"comment:角色ID;default:0"`
}
