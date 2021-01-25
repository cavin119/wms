package error

const (
	//通用错误码
	FORBID_DATA = 9901 //禁用的数据
	ERROR   = 9999 //未知错误
	//start jwt
	TOKEN_REQUIRED = 1001 //token不能为空
	TOKEN_BLACKLIST = 1002 //token被列入黑名单
	TOKEN_INVALID = 1003 //TOKEN不合法
	TOKEN_EXPIRED = 1004 //TOKEN已过期
	TOKEN_REPACED = 1005 //TOKEN已替换，在别的地方登录
	TOEKN_GET_FAIL = 1008 //获取token失败
	//end jwt
	//start sys_user
	USERNAME_OR_PASSWORD_ERROR = 1020 //用户名或者密码错误
	//end sys_user
)
