package main

import (
	"douyin/config"
	"douyin/router"
	"douyin/util/videoutil"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	initFilePath()
	r := gin.Default()
	router.InitRouter(r)
	r.Run(config.ServerPort)
}

func initFilePath() {
	systemPath := videoutil.GetCurrentAbPath()
	pos := strings.Index(systemPath, ":\\")
	if pos != -1 {
		// windows盘符
		config.TEMP_FILE_DIR = "\\tmp\\"
		config.RUNTIME_ENV = "\\"
	} else {
		// unix路径
		config.TEMP_FILE_DIR = "/tmp/"
		config.RUNTIME_ENV = "/"
	}
	exist, err := pathExists(systemPath + config.TEMP_FILE_DIR)
	if err != nil {
		log.Fatalln(err)
	}
	if !exist {
		// 创建文件夹
		err := os.Mkdir(systemPath+config.TEMP_FILE_DIR, os.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
