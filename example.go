package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

type User struct {
	Nickname string `json:"nickname"`
	Msg      string `json:"msg"`
	Sex      string `json:"sex" binding:"required"`
	Age      int `json:"age" binding:"required,gte=1,lte=120"`
	Email    string `json:"email" binding:"required,email"`
	Birthday time.Time `json:"birthday" binding:"required" time_format:"2006-01-02" time_utc:"8"`
}



func HandlePostJson(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(200, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	//c.BindJSON(&user)
	c.JSON(200, gin.H{
		"status": 200,
		"nickname": user.Nickname,
		"msg": user.Msg,
	})
}

type TEST struct {
	name []byte
}

func newTest() *TEST {
	return &TEST{
		[]byte("34234324"),
	}
}
func main() {
	//j := middlewares.NewJWT()
	//claims := jwt.MapClaims{
	//	"exp": json.Number(strconv.FormatInt(time.Now().Add(time.Hour*time.Duration(1)).Unix(), 10)),
	//	"iat": json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
	//	"UUID": uuid.NewV1(),
	//	"ID":   1,
	//	"Username": "cavin",
	//	"AuthorityId": "1",
	//	"StandardClaims": jwt.StandardClaims{
	//		NotBefore: time.Now().Unix() - 100,
	//		ExpiresAt: time.Now().Unix() + 86400,
	//		Issuer: "wms",
	//	},
	//}

}
func main1() {

	t := newTest()
	println(t)
	os.Exit(0)
	r := gin.Default()//默认连接了logger 和 recovery 中间件
	//r := gin.New()//没有中间件
	r.MaxMultipartMemory = 8 << 20

	r.GET("/", func(c *gin.Context) {
		c.String(200, "欢迎使用WMS")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, "您好，%s", name)
	})
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		msg := name + " is " + action
		c.String(200, msg)
	})
	r.GET("/welcome", func(c *gin.Context) {
		name := c.DefaultQuery("name", "Tom")
		job := c.DefaultQuery("job", "IT support")
		msg := name + " job is " + job
		c.String(200, msg)
	})
	r.POST("/sendMsg", func(c *gin.Context) {
		msg := c.PostForm("msg")
		nickname := c.DefaultPostForm("nickname", "Tom")

		c.JSON(200, gin.H{
			"status": 200,
			"message": msg,
			"nickname": nickname,
		})
	})
	r.POST("/sendJson", HandlePostJson)
	//%v按值的本来值输出 %+v在 %v 基础上，对结构体字段名和值进行展开
	//%#v 	输出 Go 语言语法格式的值
	//%T 输出 Go 语言语法格式的类型和值
	//%% 输出 % 本体
	//%b 整型以二进制方式显示
	//%o 整型以八进制方式显示
	//%d 整型以十进制方式显示
	//%x 整型以十六进制方式显示
	//%X 整型以十六进制、字母大写方式显示
	//%U Unicode 字符
	//%f 浮点数
	//%p 指针，十六进制方式显示
	r.POST("/sendMap", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.QueryMap("names")

		str := "结果是：" + fmt.Sprintf("ids: %v, names: %v", ids, names)

		c.String(200, str)
	})

	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		dist := "qwer.jpg"
		c.SaveUploadedFile(file, dist)

		c.String(200, fmt.Sprintf("%s 上传成功", file.Filename))

	})



	r.Run()
}
