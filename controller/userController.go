package controller

import (
	"douyin/response"
	"douyin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	if username == "" || password == "" {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 500,
			StatusMsg:  "缺少username/password参数",
		})
		c.Abort()
	}
	userSerice := service.NewUserService()
	response := userSerice.Login(username, password)
	c.JSON(http.StatusOK, response)
}
