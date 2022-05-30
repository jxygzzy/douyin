package db

import "douyin/config"

type FavoriteDao struct {
	ID      int `gorm:"column:id;autoIncrement"`
	VideoId int `gorm:"column:video_id"`
	UserId  int `gorm:"column:user_id"`
}

func (FavoriteDao) TableName() string {
	return config.FavoriteTableName
}
