package controller

import (
	"douyin/response"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Follow(c *gin.Context) {
	valid := validFollow(c)
	if !valid {
		return
	}
	to_user_id_query := c.Query("to_user_id")
	to_user_id, _ := strconv.ParseInt(to_user_id_query, 10, 64)
	action_type := c.Query("action_type")
	user_id, _ := c.Get("userId")
	relationService := service.NewRelationService()
	if action_type == "1" {
		resp, err := relationService.Follow(user_id.(int64), to_user_id)
		if err != nil {
			c.JSON(http.StatusOK, response.Response{
				StatusCode: 500,
				StatusMsg:  "关注失败：" + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, resp)
		return
	} else {
		resp, err := relationService.UnFollow(user_id.(int64), to_user_id)
		if err != nil {
			c.JSON(http.StatusOK, response.Response{
				StatusCode: 500,
				StatusMsg:  "取消关注失败：" + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, resp)
		return
	}
}

func FollowList(c *gin.Context) {
	user_id_query := c.Query("user_id")
	if user_id_query == "" {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 500,
			StatusMsg:  "user_id不能为空",
		})
		return
	}
	relationService := service.NewRelationService()
	user_id, _ := strconv.ParseInt(user_id_query, 10, 64)
	resp, err := relationService.FollowList(user_id)
	if err != nil {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 500,
			StatusMsg:  "获取失败：" + err.Error(),
		})
		return
	}
	if resp.UserList == nil {
		resp.UserList = &[]response.User{}
	}
	c.JSON(http.StatusOK, resp)
}

func GetFollowerList(c *gin.Context) {
	userString := c.Query("user_id")
	if userString == "" {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 500,
			StatusMsg:  "缺少user_id参数",
		})
		return
	}
	userId, _ := strconv.ParseInt(userString, 10, 64)
	relationService := service.NewRelationService()
	resp, err := relationService.GetFollowerList(userId)
	if err != nil {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 500,
			StatusMsg:  "获取失败：" + err.Error(),
		})
		return
	}
	if resp.UserList == nil {
		resp.UserList = &[]response.User{}
	}
	c.JSON(http.StatusOK, resp)
}

func validFollow(c *gin.Context) bool {
	to_user_id_query := c.Query("to_user_id")
	if to_user_id_query == "" {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 500,
			StatusMsg:  "to_user_id不能为空",
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
