package main

import (
	"fmt"
	"gin_jiaoshoujia/config"
	"gin_jiaoshoujia/models"
	"gin_jiaoshoujia/routers"
)

func main() {
	// 初始化配置
	config.InitConfig()

	// 初始化数据库
	models.InitDB()

	// 初始化路由
	app := routers.InitRouter()

	// 启动服务器
	port := config.Config.GetInt("app.port")
	err := app.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
}
