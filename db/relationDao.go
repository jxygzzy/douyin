package db

import (
	"douyin/config"
	"douyin/response"

	"gorm.io/gorm"
)

type RelationDao struct {
	ID       int64 `gorm:"column:id;autoIncrement"`
	UserId   int64 `gorm:"column:user_id"`
	ToUserId int64 `gorm:"column:to_user_id"`
}

func (RelationDao) TableName() string {
	return config.RelationTableName
}

func Follow(user_id int64, to_user_id int64) error {
	tx := DB.Begin()
	relationDao := &RelationDao{
		UserId:   user_id,
		ToUserId: to_user_id,
	}
	err := tx.Save(&relationDao)
	if err.Error != nil {
		tx.Rollback()
		return err.Error
	}
	err = tx.Model(&UserDao{}).
		Where("id = ?", user_id).
		UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1))
	if err.Error != nil {
		tx.Rollback()
		return err.Error
	}
	err = tx.Model(&UserDao{}).
		Where("id = ?", to_user_id).
		UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1))
	if err.Error != nil {
		tx.Rollback()
		return err.Error
	}
	tx.Commit()
	return nil
}

func UnFollow(user_id int64, to_user_id int64) error {
	tx := DB.Begin()
	err := tx.Where("user_id = ? and to_user_id = ?", user_id, to_user_id).Delete(&RelationDao{})
	if err.Error != nil {
		tx.Rollback()
		return err.Error
	}
	err = tx.Model(&UserDao{}).
		Where("id = ?", user_id).
		UpdateColumn("follow_count", gorm.Expr("follow_count - ?", 1))
	if err.Error != nil {
		tx.Rollback()
		return err.Error
	}
	err = tx.Model(&UserDao{}).
		Where("id = ?", to_user_id).
		UpdateColumn("follower_count", gorm.Expr("follower_count - ?", 1))
	if err.Error != nil {
		tx.Rollback()
		return err.Error
	}
	tx.Commit()
	return nil
}

func FollowList(user_id int64) (userList *[]response.User, err error) {
	dbErr := DB.Raw(`
	SELECT
		t_user.id AS id,
		t_user.NAME AS name,
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
		r1.user_id = ? 
		AND t_user.id = r1.to_user_id
	`, user_id).Scan(&userList)
	if dbErr.Error != nil {
		return nil, dbErr.Error
	}
	return userList, nil
}

func GetFollowerList(userId int64) (userList *[]response.User, err error) {
	dbErr := DB.Raw(`
	SELECT
		t_user.id AS id,
		t_user.NAME AS name,
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
		r1.to_user_id = ? 
		AND t_user.id = r1.user_id
	`, userId).Scan(&userList)
	if dbErr.Error != nil {
		return nil, dbErr.Error
	}
	return userList, nil
}
