package controller

import (
	"douyin/response"
	"douyin/service"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func GetFollowerList(c *gin.Context) {
	userString := c.Query("user_id")
	if userString == "" {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 500,
			StatusMsg:  "缺少user_id参数",
		})
		c.Abort()
	}
	userId, err := strconv.Atoi(userString)
	if err != nil {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 500,
			StatusMsg:  "user_id参数有误",
		})
		c.Abort()
	}
	relationService := service.NewRelationService()
	resp := relationService.GetFollowerList(userId)
	c.JSON(http.StatusOK, resp)
}
