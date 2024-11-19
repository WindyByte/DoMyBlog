package core

import (
	"backend/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func InitGormClient() *gorm.DB {
	// 初始化数据库
	if global.Config.MySQL.Host == "" {
		global.Log.Errorf("MySQL host is empty, skip init mysql")
		return nil
	}
	dsn := global.Config.MySQL.Dsn()
	global.Log.Infof("MySQL DSN: %v", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		global.Log.Errorf("Failed to connect database, err: %v", err)
		return nil
	}
	// 执行查询以获取当前数据库中的表列表
	var tables []string
	result := db.Raw("SHOW TABLES").Scan(&tables)
	if result.Error != nil {
		log.Fatalf("Failed to query tables: %v", result.Error)
	}
	// 输出表列表
	global.Log.Infof("Tables in database: %v", tables)
	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to get sql.MySQLClient")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}
