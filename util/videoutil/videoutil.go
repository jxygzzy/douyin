package videoutil

import (
	"bytes"
	"context"
	"douyin/config"
	"douyin/util/redisutil"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	imaging "github.com/disintegration/imaging"
	"github.com/gomodule/redigo/redis"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// 方法参照 https://juejin.cn/post/7099827417170051103
// GetSnapshot 生成视频缩略图并保存（作为封面）
func GetSnapshot(videoPath string, snapshotPath string, frameNum int) (snapshotName string, err error) {
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		panic(err)
	}

	err = imaging.Save(img, snapshotPath+".jpeg")
	if err != nil {
		panic(err)
	}

	// 成功则返回生成的缩略图名
	names := strings.Split(snapshotPath, "\\")
	snapshotName = names[len(names)-1] + ".jpeg"
	return
}

func UploadData(key string, data []byte) {
	putPolicy := storage.PutPolicy{
		Scope: config.QiniuBucket,
	}
	mac := qbox.NewMac(config.QiniuAccessKey, config.QiniuSecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	bucketManager := storage.NewBucketManager(mac, &cfg)
	fileInfo, sErr := bucketManager.Stat(config.QiniuBucket, key)
	if sErr == nil && fileInfo.Fsize != 0 {
		// 当文件在云端存在则不上传
		return
	}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	dataLen := int64(len(data))
	err := formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
}

var (
	redis_pool_config = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.Redis_addr,
				redis.DialDatabase(config.Redis_db),
				redis.DialPassword(config.Redis_password))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
)

func GetDownloadUrl(key string) string {
	ru := redisutil.NewRedisUtil(redis_pool_config)
	var url string
	hit, err := ru.Get(context.Background(), key, &url)
	if err != nil {
		log.Fatal(err)
	}
	if hit {
		return url
	}
	mac := qbox.NewMac(config.QiniuAccessKey, config.QiniuSecretKey)
	domain := config.QiniuDomian
	deadline := time.Now().Add(time.Second * config.QiniuUrlExpire).Unix() //1小时有效期
	privateAccessURL := storage.MakePrivateURL(mac, domain, key, deadline)
	// 缓存url，避免频繁创建
	go ru.Set(context.Background(), key, privateAccessURL, config.QiniuUrlExpire-config.RedisUrlExpireDiff)
	return privateAccessURL
}

func GetCurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(dir, tmpDir) {
		return getCurrentAbPathByCaller()
	}
	return dir
}

// 获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
