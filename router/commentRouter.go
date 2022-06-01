package router

import (
	"douyin/middleware"

	"github.com/gin-gonic/gin"
)

func InitCommentRouter(r *gin.Engine) {
	r.POST("/douyin/comment/action/", middleware.AuthMiddleware())
	r.GET("/douyin/comment/list/", middleware.AuthMiddleware())
}
