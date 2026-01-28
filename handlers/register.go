package handlers

import (
	"net/http"
	"realworld/models"

	"github.com/gin-gonic/gin"
)

// POST /users
func RegisterUser(c *gin.Context) {
	// 绑定请求 JSON
	var req models.NewUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": gin.H{"body": []string{"invalid request"}},
		})
		return
	}

	// 假实现：这里你可以查数据库是否已经存在用户
	email := req.User.Email
	username := req.User.Username
	// password := req.User.Password // 生产中要加密存储

	// 返回响应
	resp := models.UserResponse{
		User: models.User{
			Email:    email,
			Username: username,
			Token:    "fake-jwt-token", // 生产中生成 JWT
			Bio:      "",
			Image:    "",
		},
	}

	c.JSON(http.StatusCreated, resp)
}
