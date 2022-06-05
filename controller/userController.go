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
		return
	}
	userSerice := service.NewUserService()
	resp, err := userSerice.Login(username, password)
	if err != nil {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 500,
			StatusMsg:  "注册失败：" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func UserRegister(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	if username == "" || password == "" {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 500,
			StatusMsg:  "缺少username/password参数",
		})
		return
	}
	userSerice := service.NewUserService()
	resp, err := userSerice.Register(username, password)
	if err != nil {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 500,
			StatusMsg:  "注册失败：" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
