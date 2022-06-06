package router

import (
	"douyin/controller"
	"douyin/middleware"

	"github.com/gin-gonic/gin"
)

func InitFavotiteRouter(r *gin.Engine) {
	r.POST("/douyin/favorite/action/", middleware.AuthMiddleware(), controller.FavoriteAction)
	r.GET("/douyin/favorite/list/", middleware.AuthMiddleware(), controller.FavoriteList)
}
