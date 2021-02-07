package model

type SysMenu struct {
	WMSModel
	ParentId  int `json:"parent_id" gorm:"comment:父级菜单id"`
	Path      string `json:"path" gorm:"comment:路由path"`
	Name      string `json:"name" gorm:"comment:路由name"`
	IsHidden  bool   `json:"is_hidden" gorm:"comment:是否隐藏"`
	Component string `json:"component" gorm:"comment:前端组件路径"`
  	Sort      int    `json:"sort" gorm:"comment:排序"`
	Meta             `json:"meta" gorm:"comment:附加属性"`
	Children []SysMenu `json:"children" gorm:"-"`
}

type Meta struct {
	Title       string `json:"title" gorm:"comment:菜单名"`
	Icon        string `json:"icon" gorm:"comment:菜单图标"`
}

