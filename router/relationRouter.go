package router

import (
	"douyin/controller"
	"douyin/middleware"
	"github.com/gin-gonic/gin"
)

func InitRelationRouter(r *gin.Engine) {
	r.GET("/douyin/relation/follower/list/", middleware.AuthMiddleware(), controller.GetFollowerList)
	r.POST("/douyin/relation/action/", middleware.AuthMiddleware(), controller.Follow)
	r.GET("/douyin/relation/follow/list/", middleware.AuthMiddleware(), controller.FollowList)
}
