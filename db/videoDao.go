package db

import (
	"douyin/config"
	"log"
	"time"
)

type VideoDao struct {
	ID            int       `gorm:"column:id;autoIncrement"`
	PlayKey       string    `gorm:"column:play_key"`
	CoverKey      string    `gorm:"column:cover_key"`
	UserId        int       `gorm:"column:user_id"`
	Title         string    `gorm:"column:title"`
	CommentCount  int       `gorm:"column:comment_count"`
	FavoriteCount int       `gorm:"cloumn:favorite_count"`
	CreateDate    time.Time `gorm:"column:create_date"`
}

func (VideoDao) TableName() string {
	return config.VideoTableName
}

func SaveVideo(user_id int, play_key string, cover_key string, title string) error {
	video := &VideoDao{
		PlayKey:    play_key,
		CoverKey:   cover_key,
		UserId:     user_id,
		Title:      title,
		CreateDate: time.Now(),
	}
	DB.Save(video)
	if DB.Error != nil {
		log.Fatalln(DB.Error)
	}
	return DB.Error
}
