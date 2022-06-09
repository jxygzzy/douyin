package db

import (
	"douyin/config"
	"douyin/response"
)

type UserDao struct {
	ID              int64  `grorm:"column:id;autoIncrement"`
	Username        string `gorm:"column:username"`
	Password        string `gorm:"column:password"`
	Name            string `gorm:"column:name"`
	FollowCount     int64  `gorm:"column:follow_count"`
	FollowerCount   int64  `gorm:"column:follower_count"`
	Avatar          string `gorm:"column:avatar"`
	BackgroundImage string `gorm:"column:background_image"`
	Signature       string `gorm:"column:signature"`
}

func (UserDao) TableName() string {
	return config.UserTableName
}

func GetUserByUsername(username string) (*UserDao, error) {
	userDao := &UserDao{}
	err := DB.Where("username", username).First(&userDao)
	if err.Error != nil {
		return nil, err.Error
	}
	return userDao, nil
}

func GetAuthorById(user_id int64, author_id int64) (author response.User) {
	DB.Raw(`
	select t_user.id AS id,
	t_user.NAME AS name,
	t_user.follow_count AS follow_count,
	t_user.follower_count AS follower_count,
	t_user.avatar as avatar,
	t_user.background_image as background_image,
	t_user.signature as signature,
	(select count(*) from t_favorite,t_video 
	where t_video.user_id = t_user.id 
	and t_favorite.video_id = t_video.id) as total_favorited,
	(select count(*) from t_favorite where user_id = t_user.id ) as favorite_count,
	IF((
		SELECT
			count(*) 
		FROM
			t_relation
		WHERE
		t_relation.to_user_id = t_user.id
		AND t_relation.user_id = ?
		) > 0,
		TRUE,
		FALSE 
	) AS is_follow 
	from t_user where t_user.id = ?
	`, user_id, author_id).Scan(&author)
	return
}

func GetUserById(user_id int64, to_user_id int64) (user *response.User, err error) {
	dbErr := DB.Raw(`
	select t_user.id AS id,
	t_user.NAME AS name,
	t_user.follow_count AS follow_count,
	t_user.follower_count AS follower_count ,
	t_user.avatar as avatar,
	t_user.background_image as background_image,
	t_user.signature as signature,
	(select count(*) from t_favorite,t_video 
	where t_video.user_id = t_user.id 
	and t_favorite.video_id = t_video.id) as total_favorited,
	(select count(*) from t_favorite where user_id = t_user.id ) as favorite_count,
	IF((
		SELECT
			count(*) 
		FROM
			t_relation
		WHERE
		t_relation.to_user_id = t_user.id
		AND t_relation.user_id = ?
		) > 0,
		TRUE,
		FALSE 
	) AS is_follow 
	from t_user where t_user.id = ?
	`, user_id, to_user_id).Scan(&user)
	if dbErr.Error != nil {
		return nil, dbErr.Error
	}
	return user, nil
}

func Register(username string, password string, name string) (*UserDao, error) {
	userDao := &UserDao{
		Username: username,
		Password: password,
		Name:     name,
	}
	err := DB.Save(&userDao)
	if err.Error != nil {
		return nil, err.Error
	}
	return userDao, nil
}
