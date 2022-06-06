package service

import (
	"douyin/db"
	"douyin/response"
	"douyin/util/videoutil"
	"sort"
	"sync"
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
		StatusCode: 0,
		StatusMsg:  "点赞成功",
	}, nil
}

func (fs *FavoriteService) UnFavorite(user_id int64, video_id int64) (*response.Response, error) {
	err := db.UnFavorite(user_id, video_id)
	if err != nil {
		return nil, err
	}
	return &response.Response{
		StatusCode: 0,
		StatusMsg:  "取消点赞成功",
	}, nil
}

type FavoriteListResponse struct {
	response.Response
	VideoList *[]response.Video `json:"video_list"`
}

func (fs *FavoriteService) FavoriteList(user_id int64) (*FavoriteListResponse, error) {
	videoDaos, err := db.FavoriteList(user_id)
	if err != nil {
		return nil, err
	}
	if videoDaos == nil {
		return &FavoriteListResponse{
			Response: response.Response{
				StatusCode: 0,
				StatusMsg:  "喜欢列表为空",
			},
		}, nil
	}
	wg := sync.WaitGroup{}
	videoList := make([]response.Video, 0, len(*videoDaos))
	for i, n := 0, len(*videoDaos); i < n; i++ {
		wg.Add(1)
		go func(videoDao db.VideoDao) {
			defer wg.Done()
			var video = &response.Video{}
			video.Id = videoDao.ID
			video.FavoriteCount = videoDao.FavoriteCount
			video.CommentCount = videoDao.CommentCount
			video.Title = videoDao.Title
			video.Author = db.GetAuthorById(user_id, videoDao.UserId)
			video.IsFavorite = true
			video.PlayUrl = videoutil.GetDownloadUrl(videoDao.PlayKey)
			video.CoverUrl = videoutil.GetDownloadUrl(videoDao.CoverKey)
			video.CreateDate = videoDao.CreateDate
			videoList = append(videoList, *video)
		}((*videoDaos)[i])
	}
	wg.Wait()
	sort.Slice(videoList, func(i, j int) bool {
		return videoList[i].CreateDate.After(videoList[j].CreateDate)
	})
	return &FavoriteListResponse{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "获取点赞列表成功",
		},
		VideoList: &videoList,
	}, nil
}
