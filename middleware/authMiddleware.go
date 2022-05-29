package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token != "" {
			// 向redis查询token信息
			c.Next()
		}
		token = c.Request.PostFormValue("token")
		if token != "" {
			// 向redis查询token信息
			c.Next()
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 2003,
			"msg":  "请求头中auth为空",
		})
		c.Abort()
	}
}
