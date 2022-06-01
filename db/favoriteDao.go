package db

import (
	"douyin/config"
)

type FavoriteDao struct {
	ID      int `gorm:"column:id;autoIncrement"`
	VideoId int `gorm:"column:video_id"`
	UserId  int `gorm:"column:user_id"`
}

func (FavoriteDao) TableName() string {
	return config.FavoriteTableName
}

func HasFavorite(user_id int64, video_id int64) (hasFavorite bool) {
	var count int64
	DB.Model(&FavoriteDao{}).Where("user_id = ? and video_id = ?", user_id, video_id).Count(&count)
	return count > 0
}
