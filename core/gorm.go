package core

import (
	"backend/global"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

func InitGormClient() *gorm.DB {
	// 初始化数据库
	if global.Config.MySQL.Host == "" {
		return nil
	}
	db, err := gorm.Open(sqlite.Open(global.Config.MySQL.Database), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 自动迁移数据表
	//err = db.AutoMigrate(&models.UserInfo{})
	//if err != nil {
	//	return nil
	//}
	//db.Create(&models.UserInfo{
	//	Username: "admin",
	//	Passwd:   "admin",
	//})
	//db.Create(&models.UserInfo{
	//	Username: "admin2",
	//	Passwd:   "admin2",
	//})
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get sql.MySQLClient")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	log.Printf("db: %+v", db)
	return db
}
