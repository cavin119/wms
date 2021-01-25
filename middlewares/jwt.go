package middlewares

import (
	appError "7youo-wms/app/error"
	"7youo-wms/app/response"
	"7youo-wms/global"
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)
var ctx = context.Background()

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithCodeAndMsg(appError.TOKEN_REQUIRED, "未登录或非法访问", c)
			c.Abort()
			return
		}
		if IsBlacklist(token) {
			response.FailWithCodeAndMsg(appError.TOKEN_BLACKLIST, "您的帐户异地登陆或令牌失效", c)
			c.Abort()
			return
		}
		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				response.FailWithCodeAndMsg(appError.TOKEN_EXPIRED, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithCodeAndMsg(appError.TOKEN_INVALID, "令牌不合法", c)
			c.Abort()
			return
		}
		if (claims.Exp - time.Now().Unix()) < 0 {
			response.FailWithCodeAndMsg(appError.TOKEN_EXPIRED, "授权已过期", c)
			c.Abort()
			return
		} else if token != j.GetRedisJWT(claims.Username) {
			response.FailWithCodeAndMsg(appError.TOKEN_REPACED, "账号在别的设备已登录", c)
			c.Abort()
			return
		}
		c.Set("username", claims.Username)
		c.Set("tokenExp", claims.Exp)
		c.Next()
	}
}

type JWT struct {
	SigningKey interface{}
}

type JWTClaims struct {
	Username string `json:"username"`
	Exp int64       `json:"exp"`
	jwt.MapClaims   `json:"claims"`
}

var (
	TokenExpired     = errors.New("令牌已过期")
	TokenNotValidYet = errors.New("令牌不能使用")
	TokenMalformed   = errors.New("令牌不合法")
	TokenInvalid     = errors.New("令牌不合法")
)
//配置文件
func NewJWT() *JWT {
	s := global.G_CONFIG.JWT.SigningKey
	if s == "" {
		panic(fmt.Errorf("请检查config.yaml文件jwt配置"))
	}
	return &JWT{
		[]byte(s),
	}
}
//创建token
func (j *JWT) CreateToken(claims *JWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}
//解析token
func (j *JWT) ParseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}

func (j *JWT) Claims(username string) *JWTClaims {
	exp := time.Now().Unix() + global.G_CONFIG.JWT.ExpiresTime
	claims := &JWTClaims{
		Exp: exp,
		Username: username,
		MapClaims: jwt.MapClaims{
			"iss": global.G_CONFIG.System.Name,
		},
	}
	return claims
}

//黑名单保存的redis key
func TokenBlacklistKey() string  {
	return global.G_CONFIG.System.Name + "_token_blacklist"
}
//如果在黑名单列表
func IsBlacklist(token string) bool {
	val, err := global.G_REDIS.SIsMember(ctx, TokenBlacklistKey(), token).Result()
	if err != nil {
		global.G_LOG.Error("token blacklist error", zap.Any("err", err.Error()))
		return false
	}
	return val
}
//获取用户token的key
func GetUserTokenKey(username string) string {
	return global.G_CONFIG.System.Name + "_jwt_user_" + username
}

func (j *JWT) GetRedisJWT(username string) (redisJWT string) {
	k := GetUserTokenKey(username)
	redisJWT, _ = global.G_REDIS.Get(ctx, k).Result()
	return redisJWT
}

func (j *JWT) SetRedisJWT(username string, jwt string) (err error) {
	timer := time.Duration(global.G_CONFIG.JWT.ExpiresTime) * time.Second
	k := GetUserTokenKey(username)
	err = global.G_REDIS.Set(ctx, k, jwt, timer).Err()
	return err
}