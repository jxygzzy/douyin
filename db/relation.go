package db

import "douyin/config"

type RelationDao struct {
	ID       int `gorm:"column:id;autoIncrement"`
	UserId   int `gorm:"column:user_id"`
	ToUserId int `gorm:"column:to_user_id"`
}

func (RelationDao) TableName() string {
	return config.RelationTableName
}
