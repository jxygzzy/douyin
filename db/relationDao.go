package db

import (
	"douyin/config"
	"douyin/response"
)

type RelationDao struct {
	ID       int `gorm:"column:id;autoIncrement"`
	UserId   int `gorm:"column:user_id"`
	ToUserId int `gorm:"column:to_user_id"`
}

func (RelationDao) TableName() string {
	return config.RelationTableName
}

func GetFollowerList(userId int, userList *[]response.User) error {
	DB.Raw(`
	SELECT
		t_user.id AS id,
		t_user.NAME AS NAME,
		t_user.follow_count AS follow_count,
		t_user.follower_count AS foolower_count,
	IF
		((
			SELECT
				count(*) 
			FROM
				t_relation r2 
			WHERE
				r2.to_user_id = r1.user_id 
				AND r2.user_id = r1.to_user_id 
			) > 0,
			TRUE,
			FALSE 
		) AS is_follow 
	FROM
		t_relation r1,
		t_user 
	WHERE
		r1.to_user_id = 1 
		AND t_user.id = r1.user_id
	`, userId).Scan(&userList)
	return DB.Error
}
