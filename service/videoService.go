package service

import (
	"douyin/db"
	"douyin/response"
	"douyin/util/md5util"
	"douyin/util/videoutil"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type VideoService struct {
}

func NewVideoService() *VideoService {
	return &VideoService{}
}

func (vs *VideoService) UploadData(user_id int, title string, fileName string, filePath string) *response.Response {
	var imgKey string
	var videoKey string
	go func() {
		imgPath := strings.Replace(filePath+fileName, path.Ext(fileName), "", 1)
		imgName, err := videoutil.GetSnapshot(filePath+fileName, imgPath, 48)
		if err != nil {
			panic(err)
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
		db.SaveVideo(user_id, videoKey, imgKey, title)
		os.Remove(filePath + fileName)
		os.Remove(filePath + imgName)
	}()
	return &response.Response{
		StatusCode: 200,
		StatusMsg:  "投稿成功，正在上传",
	}
}
