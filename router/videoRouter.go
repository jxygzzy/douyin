package router

import (
	"douyin/controller"
	"douyin/middleware"
	"github.com/gin-gonic/gin"
)

func InitVideoRouter(r *gin.Engine) {
	// 在绑定handler之前绑定中间件才能保证调用顺序
	publish := r.Group("/douyin/publish/").Use(middleware.AuthMiddleware())
	publish.POST("action/", controller.PublishVideo)
}
