package db

type UserDao struct {
	ID       int    `grorm:"column:id;autoIncrement"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Name     string `gorm:"column:name"`
}

