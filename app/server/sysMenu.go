package server

import (
	"7youo-wms/app/model"
	"7youo-wms/app/request"
	"7youo-wms/global"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	MenuNotFound = errors.New("菜单不存在或者已删除")
	HaveChildMenusCannotDelete = errors.New("存在下级菜单不能删除")
)

func CreateMenu(request *request.CreateMenu) (*model.SysMenu, error){
	var menu model.SysMenu
	menu.ParentId = request.ParentId
	menu.Name = request.Name
	menu.Path = request.Path
	menu.Component = request.Component
	menu.IsHidden = *request.IsHidden
	menu.Sort = request.Sort
	menu.Meta.Title = request.Meta.Title
	menu.Meta.Icon = request.Meta.Icon

	err := global.G_DB.Create(&menu).Error
	return &menu, err

}

//通过id获取角色详情
//@param roleId int
//@return nil
//@return SysRole
func GetMenuById(menuId int) (menu *model.SysMenu, err error)  {
	var model model.SysMenu
	rows := global.G_DB.Where("id = ?", menuId).Find(&model).RowsAffected
	if rows == 0 {
		return nil, MenuNotFound
	}
	return &model, nil
}

//删除角色
func DeleteMenu(menuId int) (bool, error) {
	menu, err := GetMenuById(menuId)
	//找不到记录
	if err != nil {
		return false, MenuNotFound
	}
	//存在子角色
	if !errors.Is(global.G_DB.Where("parent_id = ?", menuId).First(&model.SysMenu{}).Error, gorm.ErrRecordNotFound) {
		return false, HaveChildMenusCannotDelete
	}

	delErr := global.G_DB.Delete(menu).Error
	//数据操作时出现错误
	if delErr != nil {
		global.G_LOG.Error("DeleteMenu error", zap.Any("err", delErr.Error()))
		return false, err
	}
	return true, nil
}

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
