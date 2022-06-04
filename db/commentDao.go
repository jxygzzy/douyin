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
	DeleteAt   int64     `gorm:"column:delete_at"`
}

func (CommentDao) TableName() string {
	return config.CommentTableName
}

func SaveComment(video_id int64, user_id int64, content string) (*CommentDao, error) {
	commentDao := &CommentDao{
		VideoId:    video_id,
		UserId:     user_id,
		Content:    content,
		CreateDate: time.Now(),
		DeleteAt:   0,
	}
	err := DB.Save(commentDao)
	if err.Error != nil {
		return nil, err.Error
	}
	return commentDao, nil
}
