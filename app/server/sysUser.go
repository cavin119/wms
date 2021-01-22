package server

import (
	"7youo-wms/app/model"
	"7youo-wms/app/request"
	"7youo-wms/global"
	"7youo-wms/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

//判断用户名是否存在
func UsernameExist(username string) bool {
	//@todo 如果redis 有缓存用户信息，可以优先查询redis
	var model model.SysUser
	//没找到
	if errors.Is(global.G_DB.Select("id").Where("username = ?", username).First(&model).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
//用户注册
func Register(user *request.Register) (regUser *model.SysUser, err error)  {
	var model model.SysUser
	model.Username = user.Username
	model.Password = utils.MD5([]byte(user.Password))
	err = global.G_DB.Create(&model).Error
	return &model, err
}

//用户登录
func Login(u request.Login) (loginUser *model.SysUser, err error)  {
	var user model.SysUser
	err = global.G_DB.Where("username = ?", u.Username).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("user_not_found")
	}
	if user.Password != utils.MD5([]byte(u.Password)) {
		return nil, errors.New("password_error")
	}
	return &user, nil
}

//获取用户信息
func GetUserInfo(username string) (user *model.SysUser, err error) {
	var model model.SysUser
	err = global.G_DB.Where("username = ?", username).First(&model).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("not_found")
	}
	fmt.Printf("get user info err %v", err)
	fmt.Printf("get user info %v", user)
	return &model, nil
}