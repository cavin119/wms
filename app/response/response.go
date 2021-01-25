package response

import (
	appError "7youo-wms/app/error"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type FailValidateResponse struct {
	Code int         `json:"code"`
	Errs interface{} `json:"errs"`
	Msg  string      `json:"msg"`
}


const (
	//通用接口状态码
	SUCCESS = 200 //成功
	BAD_REQUEST = 400 //通用数据验证失败错误
	NOT_FOUND = 404 //通用获取资源失败
	REQUEST_CONFLICT = 409 //请求与服务器目标资源冲突


)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(200, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, nil, "请求成功", c)
}

func OkWithMsg(msg string, c *gin.Context) {
	Result(SUCCESS, nil, msg, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(appError.ERROR, nil, "操作失败", c)
}

func FailWithCode(code int, c *gin.Context) {
	Result(code, nil, "请求失败", c)
}
func FailWithCodeAndMsg(code int, msg string, c *gin.Context) {
	Result(code, nil, msg, c)
}

func FailWithMsg(msg string, c *gin.Context) {
	Result(appError.ERROR, nil, msg, c)
}

func FailWithDetailed(code int, data interface{}, message string, c *gin.Context) {
	Result(code, data, message, c)
}

func FailWithValidate(errs interface{}, c *gin.Context) {
	c.JSON(200, FailValidateResponse{
		BAD_REQUEST,
		errs,
		"请求数据错误",
	})
}

func FailWithConflict(errs interface{}, c *gin.Context)  {
	c.JSON(200, FailValidateResponse{
		REQUEST_CONFLICT,
		errs,
		"请求资源冲突",
	})
}