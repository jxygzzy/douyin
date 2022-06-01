package service

import (
	"douyin/db"
	"douyin/response"
	"douyin/util/md5util"
	"douyin/util/videoutil"
	"io/ioutil"
	"log"
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
