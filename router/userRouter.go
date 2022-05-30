package router

import (
	"douyin/controller"

	"github.com/gin-gonic/gin"
)

func InitUserRuoter(r *gin.Engine) {
	usergroup := r.Group("/douyin/user")
	usergroup.POST("/login/", controller.UserLogin)
}
