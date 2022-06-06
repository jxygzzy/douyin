package db

import (
	"douyin/config"
	"fmt"

	"gorm.io/gorm"
)

type FavoriteDao struct {
	ID      int64 `gorm:"column:id;autoIncrement"`
	VideoId int64 `gorm:"column:video_id"`
	UserId  int64 `gorm:"column:user_id"`
}

func (FavoriteDao) TableName() string {
	return config.FavoriteTableName
}

func HasFavorite(user_id int64, video_id int64) (hasFavorite bool) {
	var count int64
	DB.Model(&FavoriteDao{}).Where("user_id = ? and video_id = ?", user_id, video_id).Count(&count)
	return count > 0
}

func Favorite(user_id int64, video_id int64) error {
	tx := DB.Begin()
	var count int64
	tx.Model(&FavoriteDao{}).Where("user_id = ? and video_id = ?", user_id, video_id).Count(&count)
	if count > 0 {
		tx.Rollback()
		return fmt.Errorf("已经点过赞了")
	}
	favoriteDao := &FavoriteDao{
		VideoId: video_id,
		UserId:  user_id,
	}
	err := tx.Save(&favoriteDao)
	if err.Error != nil {
		tx.Rollback()
		return err.Error
	}
	err = tx.Model(&VideoDao{}).
		Where("id = ?", video_id).
		UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1))
	if err.Error != nil {
		tx.Rollback()
		return err.Error
	}
	tx.Commit()
	return nil
}

func UnFavorite(user_id int64, video_id int64) error {
	tx := DB.Begin()
	var count int64
	tx.Model(&FavoriteDao{}).Where("user_id = ? and video_id = ?", user_id, video_id).Count(&count)
	if count == 0 {
		tx.Rollback()
		return fmt.Errorf("没有点赞记录")
	}
	err := tx.Where("user_id = ? and video_id = ?", user_id, video_id).Delete(&FavoriteDao{})
	if err.Error != nil {
		tx.Rollback()
		return err.Error
	}
	err = tx.Model(&VideoDao{}).
		Where("id = ?", video_id).
		UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1))
	if err.Error != nil {
		tx.Rollback()
		return err.Error
	}
	tx.Commit()
	return nil
}

func FavoriteList(user_id int64) (videoDaos *[]VideoDao, err error) {
	dbErr := DB.Raw(`
	select t_video.* from t_video,t_favorite
	where t_video.id=t_favorite.video_id
	and t_favorite.user_id = ?
	`, user_id).Scan(&videoDaos)
	if dbErr.Error != nil {
		return nil, dbErr.Error
	}
	return videoDaos, nil
}
