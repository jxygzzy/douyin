package router

import (
	"douyin/controller"
	"douyin/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserRuoter(r *gin.Engine) {
	r.GET("/douyin/user/", middleware.AuthMiddleware())
	r.POST("/douyin/user/login/", controller.UserLogin)
	r.POST("/douyin/user/register/", controller.UserRegister)
}
