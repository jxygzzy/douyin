package middleware

import (
	"douyin/constants"
	"douyin/response"
	"douyin/util/authutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	auth = authutil.NewAuthUtil()
)

func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token != "" {
			// 向redis查询token信息
			userId, err := auth.CheckToken(c, token)
			if err != nil {
				c.JSON(http.StatusOK, response.Response{
					StatusCode: 3001,
					StatusMsg:  constants.TOKEN_NOT_EXIST_ERROR,
				})
				c.Abort()
			}
			c.Set("userId", userId)
			c.Next()
		}
		token = c.Request.PostFormValue("token")
		if token != "" {
			// 向redis查询token信息
			userId, err := auth.CheckToken(c, token)
			if err != nil {
				c.JSON(http.StatusOK, response.Response{
					StatusCode: 3001,
					StatusMsg:  constants.TOKEN_NOT_EXIST_ERROR,
				})
				c.Abort()
			}
			c.Set("userId", userId)
			c.Next()
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 3001,
			"msg":  constants.TOKEN_NOT_EXIST_ERROR,
		})
		c.Abort()
	}
}
