package server

import (
	"7youo-wms/app/model"
	"7youo-wms/app/request"
	"7youo-wms/global"
	"7youo-wms/utils"
	"errors"
	"gorm.io/gorm"
)

var (
	UserNotFound = errors.New("用户不存在")
	PasswordError = errors.New("密码错误")
)

//判断用户名是否存在
func UsernameExist(username string) bool {
	_, err := GetUserInfo(username)
	if err != nil {
		//没找到用户
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
	//var user model.SysUser
	user, err := GetUserInfo(u.Username)
	//err = global.G_DB.Where("username = ?", u.Username).First(&user).Error
	if err != nil {
		return nil, err
	}
	if user.Password != utils.MD5([]byte(u.Password)) {
		return nil, PasswordError
	}
	return user, nil
}

//获取用户信息
func GetUserInfo(username string) (user *model.SysUser, err error) {
	var model model.SysUser
	err = global.G_DB.Where("username = ?", username).First(&model).Error
	if err == gorm.ErrRecordNotFound {
		return nil, UserNotFound
	}
	return &model, nil
}