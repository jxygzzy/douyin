package db

import (
	"douyin/config"
	"time"

	"gorm.io/gorm"
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
	tx := DB.Begin()
	err := tx.Save(commentDao)
	if err.Error != nil {
		tx.Rollback()
		return nil, err.Error
	}
	err = tx.Model(&VideoDao{}).
		Where("id = ?", video_id).
		UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1))
	if err.Error != nil {
		tx.Rollback()
		return nil, err.Error
	}
	tx.Commit()
	return commentDao, nil
}

func DelComment(video_id int64, comment_id int64) error {
	tx := DB.Begin()
	err := tx.Model(&VideoDao{}).
		Where("id = ?", video_id).
		UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1))
	if err.Error != nil {
		tx.Rollback()
		return err.Error
	}
	err = tx.Model(&CommentDao{}).
		Where("id = ?", comment_id).
		UpdateColumn("delete_at", 1)
	if err.Error != nil {
		tx.Rollback()
		return err.Error
	}
	tx.Commit()
	return nil
}

func CommentList(video_id int64) (commentList *[]CommentDao, err error) {
	dbErr := DB.Model(&CommentDao{}).Where("delete_at != 1 and video_id = ?", video_id).Order("create_date desc").Scan(&commentList)
	if dbErr.Error != nil {
		return nil, dbErr.Error
	}
	return commentList, nil
}
