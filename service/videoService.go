package service

import (
	"douyin/config"
	"douyin/db"
	"douyin/response"
	"douyin/util/md5util"
	"douyin/util/videoutil"
	"io/ioutil"
	"log"
	"os"
	"path"
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

func (vs *VideoService) Feed(user_id int64, last_time time.Time) (resp *FeedResponse) {
	var next_time time.Time
	var videoList = make([]response.Video, 0, config.FEED_NUM)
	var wg sync.WaitGroup
	videos := db.Feed(last_time)
	for i, n := 0, len(*videos); i < n; i++ {
		var video = &response.Video{}
		video.Id = (*videos)[i].ID
		video.FavoriteCount = (*videos)[i].FavoriteCount
		video.CommentCount = (*videos)[i].CommentCount
		video.Title = (*videos)[i].Title
		video.Author = db.GetAuthorById(user_id, (*videos)[i].UserId)
		video.IsFavorite = db.HasFavorite(user_id, (*videos)[i].ID)
		video.PlayUrl = videoutil.GetDownloadUrl((*videos)[i].PlayKey)
		video.CoverUrl = videoutil.GetDownloadUrl((*videos)[i].CoverKey)
		videoList = append(videoList, *video)
	}
	if len(*videos) < config.FEED_NUM {
		next_time = time.Now()
	} else {
		next_time = (*videos)[len(*videos)-1].CreateDate
	}
	wg.Wait()
	return &FeedResponse{
		Response: response.Response{
			StatusCode: 200,
			StatusMsg:  "成功",
		},
		NextTime:  next_time.Unix(),
		VideoList: &videoList,
	}
}
