package controller

import (
	"douyin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	userSerice := service.NewUserService()
	response := userSerice.Login(username, password)
	c.JSON(http.StatusOK, response)
}
