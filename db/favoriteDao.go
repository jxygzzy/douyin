package db

type FavoriteDao struct {
	ID      int `gorm:"column:id;autoIncrement"`
	VideoId int `gorm:"column:video_id"`
	UserId  int `gorm:"column:user_id"`
}
