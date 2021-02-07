package server

import (
	"7youo-wms/app/model"
	"7youo-wms/global"
	"fmt"
)

func GetMenuTree() (err error, menus []model.SysMenu) {
	err, treeMap := getAllMenus()
	menus = treeMap[0]
	for i := 0; i < len(menus); i++ {
		err = getChildren(&menus[i], treeMap)
	}
	return err, menus
}

func getChildren(menu *model.SysMenu, treeMap map[int][]model.SysMenu) (err error) {
	menu.Children = treeMap[menu.WMSModel.Id]
	for i := 0; i < len(menu.Children); i++ {
		err = getChildren(&menu.Children[i], treeMap)
	}
	return err
}

func getAllMenus() (err error, treeMap map[int][]model.SysMenu)  {
	var allMenus []model.SysMenu
	treeMap = make(map[int][]model.SysMenu)

	err = global.G_DB.Order("sort").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	fmt.Printf("getAllMenus： %v", treeMap)
	return err, treeMap
}


func AddMenu() {
	var menus = []model.SysMenu{
		{
			WMSModel: model.WMSModel{
				Id: 1,
			},
			ParentId:  0,
			Path:      "dashboard",
			Name:      "dashboard",
			IsHidden:  false,
			Component: "components/dashboard/Dashboard.vue",
			Sort:      1,
			Meta: model.Meta{
				Title: "仪表盘",
				Icon:  "s-home",
			},
		},
		{
			WMSModel: model.WMSModel{
				Id: 2,
			},
			ParentId:  0,
			Path:      "system",
			Name:      "system",
			IsHidden:  false,
			Component: "components/system/System.vue",
			Sort:      2,
			Meta: model.Meta{
				Title: "超级管理员",
				Icon:  "user-solid",
			},
		},
		{
			WMSModel: model.WMSModel{
				Id: 3,
			},
			ParentId:  2,
			Path:      "roles",
			Name:      "roles",
			IsHidden:  false,
			Component: "components/system/roles/Roles.vue",
			Sort:      3,
			Meta: model.Meta{
				Title: "角色管理",
				Icon:  "coordinate",
			},
		},
		{
			WMSModel: model.WMSModel{
				Id: 4,
			},
			ParentId:  2,
			Path:      "admin",
			Name:      "admin",
			IsHidden:  false,
			Component: "components/system/admin/Admin.vue",
			Sort:      4,
			Meta: model.Meta{
				Title: "管理员",
				Icon:  "s-custom",
			},
		},
	}
	err := global.G_DB.Create(&menus).Error
	fmt.Printf("Add menu err： %v", err)
}
