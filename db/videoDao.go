package db

import (
	"douyin/config"
	"time"
)

type VideoDao struct {
	ID            int64     `gorm:"column:id;autoIncrement"`
	PlayKey       string    `gorm:"column:play_key"`
	CoverKey      string    `gorm:"column:cover_key"`
	UserId        int64     `gorm:"column:user_id"`
	Title         string    `gorm:"column:title"`
	CommentCount  int64     `gorm:"column:comment_count"`
	FavoriteCount int64     `gorm:"cloumn:favorite_count"`
	CreateDate    time.Time `gorm:"column:create_date"`
}

func (VideoDao) TableName() string {
	return config.VideoTableName
}

func SaveVideo(user_id int64, play_key string, cover_key string, title string) error {
	video := &VideoDao{
		PlayKey:    play_key,
		CoverKey:   cover_key,
		UserId:     user_id,
		Title:      title,
		CreateDate: time.Now(),
	}
	err := DB.Save(video)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func Feed(last_time time.Time) (video_list *[]VideoDao, err error) {
	dbErr := DB.Where("create_date <= ?", last_time.Format("2006-01-02 15:04:05")).Order("create_date desc").Limit(config.FEED_NUM).Find(&video_list)
	if dbErr.Error != nil {
		return nil, dbErr.Error
	}
	return video_list, nil
}
