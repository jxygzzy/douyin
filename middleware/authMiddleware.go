package middleware

import (
	"douyin/constants"
	"douyin/response"
	"douyin/util/authutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		auth := authutil.NewAuthUtil()
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
			return
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
			return
		}
		token = c.PostForm("token")
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
			return
		}
		if token == "" {
			c.JSON(http.StatusOK, response.Response{
				StatusCode: 3001,
				StatusMsg:  constants.TOKEN_NOT_EXIST_ERROR,
			})
			c.Abort()
		}
	}
}
