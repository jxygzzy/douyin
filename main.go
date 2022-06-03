package main

import (
	"douyin/config"
	"douyin/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	r.Run(config.ServerPort)
}
