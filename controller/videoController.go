package controller

import (
	"context"
	"douyin/config"
	"douyin/response"
	"douyin/service"
	"douyin/util/authutil"
	"douyin/util/videoutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func PublishVideo(c *gin.Context) {
	user_id, _ := c.Get("userId")
	token := c.PostForm("token")
	user_id_int := user_id.(int64)
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

func Feed(c *gin.Context) {
	var latest_time time.Time
	var user_id int64
	last_time_query := c.Query("latest_time")
	if last_time_query == "" || last_time_query == "0" {
		latest_time = time.Now()
	} else {
		last_time_unix, err := strconv.ParseInt(last_time_query, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, &response.Response{
				StatusCode: 500,
				StatusMsg:  "last_time时间戳错误",
			})
			return
		}
		latestTime := time.Unix(last_time_unix/1000, 0)
		latest_time = latestTime
	}
	token := c.Query("token")
	if token != "" {
		userId, err := authutil.NewAuthUtil().CheckToken(context.Background(), token)
		if err != nil {
			c.JSON(http.StatusOK, &response.Response{
				StatusCode: 500,
				StatusMsg:  "token不存在",
			})
			return
		}
		user_id = userId
	}
	videoService := service.NewVideoService()
	resp, err := videoService.Feed(user_id, latest_time)
	if err != nil {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 500,
			StatusMsg:  "刷新视频失败",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
