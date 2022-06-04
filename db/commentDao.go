package db

import (
	"douyin/config"
	"time"
)

type CommentDao struct {
	ID         int64     `gorm:"column:id;autoIncrement"`
	VideoId    int64     `gorm:"column:video_id"`
	UserId     int64     `gorm:"column:user_id"`
	Content    string    `gorm:"column:content"`
	CreateDate time.Time `gorm:"column:create_date"`
	Delete     int64     `gorm:"column:delete"`
}

func (CommentDao) TableName() string {
	return config.CommentTableName
}
