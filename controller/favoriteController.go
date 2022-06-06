package controller

import (
	"douyin/response"
	"douyin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FavoriteAction(c *gin.Context) {
	if !validFavorite(c) {
		return
	}
	user_id, _ := c.Get("userId")
	video_id, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	action_type := c.Query("action_type")
	favoriteService := service.NewFavoriteService()
	if action_type == "1" {
		resp, err := favoriteService.Favorite(user_id.(int64), video_id)
		if err != nil {
			c.JSON(http.StatusOK, response.Response{
				StatusCode: 500,
				StatusMsg:  "点赞失败：" + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, resp)
		return
	} else {
		resp, err := favoriteService.UnFavorite(user_id.(int64), video_id)
		if err != nil {
			c.JSON(http.StatusOK, response.Response{
				StatusCode: 500,
				StatusMsg:  "取消点赞失败：" + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, resp)
		return
	}
}

func validFavorite(c *gin.Context) bool {
	video_id_query := c.Query("video_id")
	if video_id_query == "" {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 500,
			StatusMsg:  "video_id不能为空",
		})
		return false
	}
	action_type := c.Query("action_type")
	if action_type != "1" && action_type != "2" {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 500,
			StatusMsg:  "action_type不能为空",
		})
		return false
	}
	return true
}

func FavoriteList(c *gin.Context) {

}
