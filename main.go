package main

import (
	"7youo-wms/bootstrap"
	"7youo-wms/global"
)

func main() {
	global.G_VP = bootstrap.Viper("./config.yaml")//加载配置文件
	global.G_LOG = bootstrap.Zap()//日志
	global.G_REDIS = bootstrap.Redis()//redis
	global.G_DB = bootstrap.Gorm() //初始化数据库连接
	db, _ := global.G_DB.DB()
	defer db.Close()

	router := bootstrap.Router()
	bootstrap.Server(global.G_CONFIG.System.Addr, router)
}
