package db

import (
	"douyin/config"
	"time"
)

type CommentDao struct {
	ID         int       `gorm:"column:id;autoIncrement"`
	VideoId    int       `gorm:"column:video_id"`
	UserId     int       `gorm:"column:user_id"`
	Content    string    `gorm:"column:content"`
	CreateDate time.Time `gorm:"column:create_date"`
}

func (CommentDao) TableName() string {
	return config.CommentTableName
}
