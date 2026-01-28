package handlers

import (
	"net/http"

	"realworld/models"

	"github.com/gin-gonic/gin"
)

// POST /users/login
func Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(422, gin.H{
			"errors": gin.H{
				"body": []string{"invalid request"},
			},
		})
		return
	}

	// 假实现：你可以自己写查数据库
	email := req.User.Email

	// 返回假数据
	var resp models.UserResponse
	resp.User.Email = email
	resp.User.Username = "fakeuser"
	resp.User.Token = "fake-jwt-token"
	resp.User.Bio = "this is a bio"
	resp.User.Image = "https://placekitten.com/200/200"

	c.JSON(http.StatusOK, resp)
}
