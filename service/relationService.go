package service

import (
	"douyin/db"
	"douyin/response"
)

type RelationService struct {
}

type FollowerListResponse struct {
	response.Response
	UserList []response.User `json:"user_list"`
}

func NewRelationService() *RelationService {
	return &RelationService{}
}

func (rs *RelationService) GetFollowerList(userId int) (resp *FollowerListResponse) {
	var userList *[]response.User
	err := db.GetFollowerList(userId, userList)
	if err != nil {
		return &FollowerListResponse{
			Response: response.Response{
				StatusCode: 500,
				StatusMsg:  "服务器错误",
			},
		}
	}
	return &FollowerListResponse{
		Response: response.Response{
			StatusCode: 200,
			StatusMsg:  "查询成功",
		},
		UserList: *userList,
	}
}
