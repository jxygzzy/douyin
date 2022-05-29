package db

import (
	"douyin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	DB *gorm.DB
)

func init() {
	db, err := gorm.Open(mysql.Open(config.DSN), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true, // 禁用默认事务
	})
	if err != nil {
		panic(err)
	}
	log.Printf("gorm-DSN:%s", config.DSN)
	DB = db
}
