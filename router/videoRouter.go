package router

import (
	"douyin/controller"
	"douyin/middleware"
	"github.com/gin-gonic/gin"
)

func InitVideoRouter(r *gin.Engine) {
	// 在绑定handler之前绑定中间件才能保证调用顺序
	r.POST("/douyin/publish/action/", middleware.AuthMiddleware(), controller.PublishVideo)
	r.GET("/douyin/feed", controller.Feed)
	r.GET("/douyin/publish/list/", middleware.AuthMiddleware(), controller.PublishList)
}
