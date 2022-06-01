package controller

import (
	"douyin/config"
	"douyin/response"
	"douyin/service"
	"douyin/util/videoutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PublishVideo(c *gin.Context) {
	user_id, exi := c.Get("userId")
	token := c.PostForm("token")
	if !exi {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 500,
			StatusMsg:  "token有误,请重新登录",
		})
		return
	}
	user_id_int := user_id.(int)
	title := c.PostForm("title")
	if title == "" {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 500,
			StatusMsg:  "title不能为空",
		})
		return
	}
	header, _ := c.FormFile("data")
	if header == nil {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 500,
			StatusMsg:  "data不能为空",
		})
		return
	}
	filePath := videoutil.GetCurrentAbPath() + config.TEMP_FILE_DIR
	header.Filename = token + header.Filename
	err := c.SaveUploadedFile(header, filePath+header.Filename)
	if err != nil {
		log.Fatalln(err)
	}
	vs := service.NewVideoService()
	resp := vs.UploadData(user_id_int, title, header.Filename, filePath)
	c.JSON(http.StatusOK, resp)
}
