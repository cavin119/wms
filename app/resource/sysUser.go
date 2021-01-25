package resource

import (
	appError "7youo-wms/app/error"
	"7youo-wms/app/model"
	"7youo-wms/app/request"
	"7youo-wms/app/response"
	"7youo-wms/app/server"
	"7youo-wms/global"
	"7youo-wms/middlewares"
	"7youo-wms/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//登录，获取token
func Login(c *gin.Context) {
	var request request.Login
	c.ShouldBindJSON(&request)
	errsMap, ok := response.RequestValidate(&request)
	if !ok {
		response.FailWithValidate(errsMap, c)
		return
	}
	user, err := server.Login(request)
	if err != nil {
		//不能提示用户名不存在
		response.FailWithCodeAndMsg(appError.USERNAME_OR_PASSWORD_ERROR, "用户名或者密码错误", c)
		return
	}
	getToken(user, c)
}
//登录用，获取token
func getToken(user *model.SysUser, c *gin.Context) {
	j := middlewares.NewJWT()
	var username = user.Username
	claims := j.Claims(username)
	tokenString, err := j.CreateToken(claims)
	if err != nil {
		response.FailWithCodeAndMsg(appError.TOEKN_GET_FAIL, "授权失败，请检查系统配置", c)
		return
	}
	j.SetRedisJWT(username, tokenString)
	response.OkWithDetailed(gin.H{
		"username": username,
		"token": tokenString,
		"expires_at": utils.UnixToTime(claims.Exp, ""),
	}, "登录成功", c)
}


//用户注册 27534
func Register(c *gin.Context) {
	var request request.Register
	c.ShouldBindJSON(&request)
	//验证请求数据
	errsMap, ok := response.RequestValidate(&request)
	if !ok {
		response.FailWithValidate(errsMap, c)
		return
	}
	//判断用户名是否被注册
	if server.UsernameExist(request.Username) {
		response.FailWithConflict(gin.H{"username": "用户名已注册"}, c)
		return
	}
	//调用注册服务
	regUser, err := server.Register(&request)
	if err != nil {
		//插入数据失败
		global.G_LOG.Error("[Register]注册失败，插入用户数据失败：", zap.Any("err", err.Error()))
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{
		"username": regUser.Username,
	}, "注册成功", c)
}
//获取用户信息
func UserInfo(c *gin.Context) {
	//获取jwt 传入的的username值
	username, _ := c.Get("username")
	var usernameStr string = username.(string)

	info, err := server.GetUserInfo(usernameStr)
	if err != nil {
		response.FailWithCodeAndMsg(response.NOT_FOUND, "用户不存在", c)
		return
	}
	response.OkWithDetailed(info, "用户信息", c)
}
