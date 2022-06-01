package router

import (
	"douyin/middleware"

	"github.com/gin-gonic/gin"
)

func InitFavotiteRouter(r *gin.Engine) {
	r.POST("/douyin/favorite/action/", middleware.AuthMiddleware())
	r.GET("/douyin/favorite/list/", middleware.AuthMiddleware())
}
