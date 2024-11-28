package core

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// initGormClient 初始化数据库连接
func initGormClient() *gorm.DB {
	if GetBasicConfig().MySQL.Host == "" {
		log.Errorf("MySQL host is empty, skip init mysql")
		return nil
	}
	dsn := GetBasicConfig().MySQL.Dsn()
	GetLogger().Infof("MySQL DSN: %v", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		GetLogger().Errorf("Failed to connect database, err: %v", err)
		return nil
	}
	return db
}
