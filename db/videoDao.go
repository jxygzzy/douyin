package db

import (
	"douyin/config"
	"time"
)

type VideoDao struct {
	ID         int       `gorm:"column:id;autoIncrement"`
	PlayKey    string    `gorm:"column:play_key"`
	CoverKey   string    `gorm:"column:cover_key"`
	UserId     int       `gorm:"column:user_id"`
	Title      string    `gorm:"column:title"`
	CreateDate time.Time `gorm:"column:create_date"`
}

func (VideoDao) TableName() string {
	return config.VideoTableName
}
