package router

import (
	"douyin/controller"
	"douyin/middleware"
	"github.com/gin-gonic/gin"
)

func InitRelationRouter(r *gin.Engine) {
	relationGroup := r.Group("/douyin/relation").Use(middleware.AuthMiddleware())
	relationGroup.GET("/follower/list/", controller.GetFollowerList)
}
