package service

import (
	"douyin/config"
	"douyin/db"
	"douyin/response"
	"douyin/util/md5util"
	"douyin/util/videoutil"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"strings"
	"sync"
	"time"
)

type VideoService struct {
}

func NewVideoService() *VideoService {
	return &VideoService{}
}

func (vs *VideoService) UploadData(user_id int64, title string, fileName string, filePath string) *response.Response {
	var imgKey string
	var videoKey string
	go func() {
		imgPath := strings.Replace(filePath+fileName, path.Ext(fileName), "", 1)
		imgName, err := videoutil.GetSnapshot(filePath+fileName, imgPath, 48)
		if err != nil {
			os.Remove(filePath + imgName)
			return
		}
		imgMd5, _ := md5util.CalcFileMD5(filePath + imgName)
		imgKey = imgMd5 + path.Ext(imgName)
		imgBytes, _ := ioutil.ReadFile(filePath + imgName)
		videoutil.UploadData(imgKey, imgBytes)
		md5, err := md5util.CalcFileMD5(filePath + fileName)
		if err != nil {
			os.Remove(filePath + fileName)
			return
		}
		videoKey = md5 + path.Ext(fileName)
		videoBytes, _ := ioutil.ReadFile(filePath + fileName)
		videoutil.UploadData(videoKey, videoBytes)
		log.Printf("用户：%v 投稿文件已上传\n", user_id)
		db.SaveVideo(user_id, videoKey, imgKey, title)
		log.Printf("用户：%v 投稿数据已保存\n", user_id)
		os.Remove(filePath + fileName)
		os.Remove(filePath + imgName)
		log.Printf("用户：%v 临时文件已删除\n", user_id)
	}()
	return &response.Response{
		StatusCode: 200,
		StatusMsg:  "投稿成功，正在上传",
	}
}

type FeedResponse struct {
	response.Response
	NextTime  int64             `json:"next_time"`
	VideoList *[]response.Video `json:"video_list"`
}

func (vs *VideoService) Feed(user_id int64, last_time time.Time) (resp *FeedResponse, err error) {
	var next_time time.Time
	var videoList = make([]response.Video, 0, config.FEED_NUM)
	videos, err := db.Feed(last_time)
	if err != nil {
		return nil, err
	}
	wg := sync.WaitGroup{}
	for i, n := 0, len(*videos); i < n; i++ {
		var videoDao = (*videos)[i]
		wg.Add(1)
		go func(videoDao db.VideoDao) {
			defer wg.Done()
			var video = &response.Video{}
			video.Id = videoDao.ID
			video.FavoriteCount = videoDao.FavoriteCount
			video.CommentCount = videoDao.CommentCount
			video.Title = videoDao.Title
			video.Author = db.GetAuthorById(user_id, videoDao.UserId)
			video.IsFavorite = db.HasFavorite(user_id, videoDao.ID)
			video.PlayUrl = videoutil.GetDownloadUrl(videoDao.PlayKey)
			video.CoverUrl = videoutil.GetDownloadUrl(videoDao.CoverKey)
			video.CreateDate = videoDao.CreateDate
			videoList = append(videoList, *video)
		}(videoDao)
	}
	if len(*videos) < config.FEED_NUM {
		next_time = time.Now()
	} else {
		next_time = (*videos)[len(*videos)-1].CreateDate
	}
	wg.Wait()
	sort.Slice(videoList, func(i, j int) bool {
		return videoList[i].CreateDate.After(videoList[j].CreateDate)
	})
	return &FeedResponse{
		Response: response.Response{
			StatusCode: 200,
			StatusMsg:  fmt.Sprintf("刷新%d条视频", len(*videos)),
		},
		NextTime:  next_time.Unix(),
		VideoList: &videoList,
	}, nil
}

type PublishListResponse struct {
	response.Response
	VideoList *[]response.Video `json:"video_list"`
}

func (vs *VideoService) PublishList(user_id int64) (*PublishListResponse, error) {
	videoDaos, err := db.PublishList(user_id)
	if err != nil {
		return nil, err
	}
	user, err := db.GetUserById(user_id, user_id)
	if err != nil {
		return nil, err
	}
	videoList := make([]response.Video, 0, len(*videoDaos))
	wg := sync.WaitGroup{}
	for i, n := 0, len(*videoDaos); i < n; i++ {
		wg.Add(1)
		go func(videoDao db.VideoDao) {
			defer wg.Done()
			video := &response.Video{}
			video.Id = videoDao.ID
			video.FavoriteCount = videoDao.FavoriteCount
			video.CommentCount = videoDao.CommentCount
			video.Title = videoDao.Title
			video.Author = *user
			video.IsFavorite = db.HasFavorite(user_id, videoDao.ID)
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
	return &PublishListResponse{
		Response: response.Response{
			StatusCode: 200,
			StatusMsg:  fmt.Sprintf("共%d条发布", len(*videoDaos)),
		},
		VideoList: &videoList,
	}, nil
}
