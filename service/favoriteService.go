package service

import (
	"douyin/db"
	"douyin/response"
)

type FavoriteService struct{}

func NewFavoriteService() *FavoriteService {
	return &FavoriteService{}
}

func (fs *FavoriteService) Favorite(user_id int64, video_id int64) (*response.Response, error) {
	err := db.Favorite(user_id, video_id)
	if err != nil {
		return nil, err
	}
	return &response.Response{
		StatusCode: 200,
		StatusMsg:  "点赞成功",
	}, nil
}

func (fs *FavoriteService) UnFavorite(user_id int64, video_id int64) (*response.Response, error) {
	err := db.UnFavorite(user_id, video_id)
	if err != nil {
		return nil, err
	}
	return &response.Response{
		StatusCode: 200,
		StatusMsg:  "取消点赞成功",
	}, nil
}
