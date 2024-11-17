package core

import (
	"backend/global"
	"backend/models"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

func InitGorm() *gorm.DB {
	// 初始化数据库
	if global.Config.MySQL.Host == "" {
		return nil
	}
	db, err := gorm.Open(sqlite.Open("backend_core.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 自动迁移数据表
	err = db.AutoMigrate(&models.UserInfo{})
	if err != nil {
		return nil
	}
	db.Create(&models.UserInfo{
		Username: "admin",
		Passwd:   "admin",
	})
	db.Create(&models.UserInfo{
		Username: "admin2",
		Passwd:   "admin2",
	})
	var tables []string
	db = db.Debug()
	db.Raw("SELECT name FROM sqlite_master WHERE type='table'").Scan(&tables)
	fmt.Println("Tables in database:", tables)
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get sql.DB")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	log.Printf("db: %+v", db)
	return db
}
