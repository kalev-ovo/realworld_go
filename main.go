package main

import (
	"net/http"
	"realworld/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 登录接口
	r.POST("/users/login", handlers.Login)
	// 注册接口
	r.POST("/users", handlers.RegisterUser)

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// 启动服务
	r.Run(":3000")
}
