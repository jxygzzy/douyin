package db

import "douyin/config"

type UserDao struct {
	ID            int    `grorm:"column:id;autoIncrement"`
	Username      string `gorm:"column:username"`
	Password      string `gorm:"column:password"`
	Name          string `gorm:"column:name"`
	FollowCount   int    `gorm:"column:follow_count"`
	FollowerCount int    `gorm:"column:follower_count"`
}

func (UserDao) TableName() string {
	return config.UserTableName
}

func GetUserByUsername(username string) *UserDao {
	userDao := &UserDao{}
	DB.Where("username", username).First(&userDao)
	return userDao
}
