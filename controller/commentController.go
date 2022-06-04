package controller

import (
	"douyin/response"
	"douyin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Comment(c *gin.Context) {
	valid := validComment(c)
	if !valid {
		return
	}
	video_id_query := c.Query("video_id")
	video_id, _ := strconv.ParseInt(video_id_query, 10, 64)
	action_type_query := c.Query("action_type")
	comment_text := c.Query("comment_text")
	comment_id_query := c.Query("comment_id")
	comment_id, _ := strconv.ParseInt(comment_id_query, 10, 64)
	commentService := service.NewCommentService()
	if action_type_query == "1" {
		user_id, _ := c.Get("userId")
		resp, err := commentService.PublishComment(video_id, user_id.(int64), comment_text)
		if err != nil {
			c.JSON(http.StatusOK, response.Response{
				StatusCode: 500,
				StatusMsg:  "评论失败",
			})
			return
		}
		c.JSON(http.StatusOK, resp)
		return
	} else {
		commentService.DelComment(video_id, comment_id)
	}
}

func CommentList(c *gin.Context) {

}

func validComment(c *gin.Context) bool {
	video_id_query := c.Query("video_id")
	if video_id_query == "" {
		c.JSON(http.StatusOK, &response.Response{
			StatusCode: 500,
			StatusMsg:  "video_id不能为空",
		})
		return false
	}
	action_type_query := c.Query("action_type")
	if action_type_query != "1" && action_type_query != "2" {
		c.JSON(http.StatusOK, &response.Response{
			StatusCode: 500,
			StatusMsg:  "action_type不能为空",
		})
		return false
	}
	if comment_text := c.Query("comment_text"); action_type_query == "1" && comment_text == "" {
		c.JSON(http.StatusOK, &response.Response{
			StatusCode: 500,
			StatusMsg:  "comment_text不能为空",
		})
		return false
	}
	if comment_id := c.Query("comment_id"); action_type_query == "2" && comment_id == "" {
		c.JSON(http.StatusOK, &response.Response{
			StatusCode: 500,
			StatusMsg:  "comment_id不能为空",
		})
		return false
	}
	return true
}
