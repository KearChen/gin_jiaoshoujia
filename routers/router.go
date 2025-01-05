package routers

import (
	"gin_jiaoshoujia/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// 示例路由
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 用户模块路由
	userGroup := router.Group("/users")
	{
		userGroup.POST("/", controllers.CreateUser) // 创建用户
		userGroup.GET("/", controllers.GetUsers)    // 获取用户列表
	}

	return router
}
