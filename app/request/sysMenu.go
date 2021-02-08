package request


type CreateMenu struct {
	ParentId int `json:"parent_id" validate:"gte=0" label:"上级菜单" gorm:"default:0"`
	Name string `json:"name" validate:"required,min=3,max=20" label:"路由Name"`
	Path string `json:"path" validate:"required,min=3,max=20" label:"路由Path"`
	Component string `json:"component" validate:"required" label:"component路径"`
	IsHidden bool `json:"is_hidden" validate:"required" label:"是否隐藏"`
	Sort int `json:"sort" validate:"required" label:"排序" gorm:"default:0"`
	Meta `json:"meta"`
}

type Meta struct {
	Title string `json:"title" validate:"required,min=3,max=20" label:"菜单名称"`
	Icon string `json:"icon" validate:"required" label:"菜单图标"`
}
