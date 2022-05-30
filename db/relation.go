package db

type RelationDao struct {
	ID       int `gorm:"column:id;autoIncrement"`
	UserId   int `gorm:"column:user_id"`
	ToUserId int `gorm:"column:to_user_id"`
}
