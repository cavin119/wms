package model

type SysDict struct {
	WMSModel
	Name      string `json:"name" gorm:"comment:字典标题"`
	DictType  string `json:"dict_type" gorm:"comment:字典类型"`
	Status    bool   `json:"status" gorm:"comment:状态，开启，关闭"`
	Desc      string `json:"desc" gorm:"comment:描述"`
}


