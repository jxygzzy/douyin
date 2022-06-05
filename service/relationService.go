package service

import (
	"douyin/db"
	"douyin/response"
	"fmt"
)

type RelationService struct {
}

type FollowerListResponse struct {
	response.Response
	UserList *[]response.User `json:"user_list"`
}

func NewRelationService() *RelationService {
	return &RelationService{}
}

func (rs *RelationService) Follow(user_id int64, to_user_id int64) (*response.Response, error) {
	if user_id == to_user_id {
		return nil, fmt.Errorf("不能关注自己")
	}
	err := db.Follow(user_id, to_user_id)
	if err != nil {
		return nil, err
	}
	return &response.Response{
		StatusCode: 200,
		StatusMsg:  "关注成功",
	}, nil
}

func (rs *RelationService) UnFollow(user_id int64, to_user_id int64) (*response.Response, error) {
	if user_id == to_user_id {
		return nil, fmt.Errorf("不能取消关注自己")
	}
	err := db.UnFollow(user_id, to_user_id)
	if err != nil {
		return nil, err
	}
	return &response.Response{
		StatusCode: 200,
		StatusMsg:  "取消关注成功",
	}, nil
}

type FollowListResponse struct {
	response.Response
	UserList *[]response.User `json:"user_list"`
}

func (rs *RelationService) FollowList(user_id int64) (*FollowListResponse, error) {
	userList, err := db.FollowList(user_id)
	if err != nil {
		return nil, err
	}
	return &FollowListResponse{
		Response: response.Response{
			StatusCode: 200,
			StatusMsg:  "获取成功",
		},
		UserList: userList,
	}, nil
}

func (rs *RelationService) GetFollowerList(userId int64) (*FollowerListResponse, error) {
	userList, err := db.GetFollowerList(userId)
	if err != nil {
		return nil, err
	}
	return &FollowerListResponse{
		Response: response.Response{
			StatusCode: 200,
			StatusMsg:  "查询成功",
		},
		UserList: userList,
	}, nil
}
