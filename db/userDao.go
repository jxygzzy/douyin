package db

import (
	"douyin/config"
	"douyin/response"
)

type UserDao struct {
	ID            int64  `grorm:"column:id;autoIncrement"`
	Username      string `gorm:"column:username"`
	Password      string `gorm:"column:password"`
	Name          string `gorm:"column:name"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"column:follower_count"`
}

func (UserDao) TableName() string {
	return config.UserTableName
}

func GetUserByUsername(username string) *UserDao {
	userDao := &UserDao{}
	DB.Where("username", username).First(&userDao)
	return userDao
}

func GetAuthorById(user_id int64, author_id int64) (author response.User) {
	DB.Raw(`
	select t_user.id AS id,
	t_user.NAME AS name,
	t_user.follow_count AS follow_count,
	t_user.follower_count AS foolower_count,
	IF((
		SELECT
			count(*) 
		FROM
			t_relation
		WHERE
		t_relation.to_user_id = ?
		AND t_relation.user_id = t_user.id
		) > 0,
		TRUE,
		FALSE 
	) AS is_follow 
	from t_user where t_user.id = ?
	`, author_id, user_id).Scan(&author)
	return
}
