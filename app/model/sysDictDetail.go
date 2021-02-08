package model

type SysDictDetail struct {
	WMSModel
	Label  string `json:"label" gorm:"comment:数据字典值名"`
	Value  int `json:"value" gorm:"comment:数据字典值"`
	Sort   int   `json:"sort" gorm:"comment:排序"`
	Desc   string `json:"desc" gorm:"comment:描述"`
	DictId int `json:"dict_id" gorm:"comment:关联的主表id"`
}


