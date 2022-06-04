package router

import (
	"douyin/controller"
	"douyin/middleware"

	"github.com/gin-gonic/gin"
)

func InitCommentRouter(r *gin.Engine) {
	r.POST("/douyin/comment/action/", middleware.AuthMiddleware(), controller.Comment)
	r.GET("/douyin/comment/list/", middleware.AuthMiddleware(), controller.CommentList)
}
