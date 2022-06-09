package db

import (
	"douyin/config"
	"douyin/response"
	"fmt"

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
	var count int64
	tx.Model(&RelationDao{}).Where("user_id = ? and to_user_id = ?", user_id, to_user_id).Count(&count)
	if count > 0 {
		tx.Rollback()
		return fmt.Errorf("已经关注了")
	}
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
	var count int64
	tx.Model(&RelationDao{}).Where("user_id = ? and to_user_id = ?", user_id, to_user_id).Count(&count)
	if count == 0 {
		tx.Rollback()
		return fmt.Errorf("没有关注记录")
	}
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
		t_user.avatar as avatar,
		t_user.background_image as background_image,
		t_user.signature as signature,
		(select count(*) from t_favorite,t_video 
		where t_video.user_id = t_user.id 
		and t_favorite.video_id = t_video.id) as total_favorited,
		(select count(*) from t_favorite where user_id = t_user.id ) as favorite_count,
		'true' is_follow 
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
		t_user.avatar as avatar,
		t_user.background_image as background_image,
		t_user.signature as signature,
		(select count(*) from t_favorite,t_video 
		where t_video.user_id = t_user.id 
		and t_favorite.video_id = t_video.id) as total_favorited,
		(select count(*) from t_favorite where user_id = t_user.id ) as favorite_count,
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
