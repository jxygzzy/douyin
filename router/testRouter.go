package router

import (
	"douyin/controller"
	"github.com/gin-gonic/gin"
)

func initTestRuoter(r *gin.Engine) {
	testRuoter := r.Group("/test")
	testRuoter.GET("/", controller.Hello)
}
