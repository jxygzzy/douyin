package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "hello world")
}
