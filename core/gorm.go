package core

import (
	"backend/model"
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
	for _, table := range GetBasicConfig().MySQL.Tables {
		if table == "" {
			continue
		}
		if db.Migrator().HasTable(table) {
			continue
		}
		GetLogger().Infof("AutoMigrate Table: %v", table)
		switch table {
		case "user_infos":
			err = db.AutoMigrate(&model.UserInfo{})
			if err != nil {
				GetLogger().Errorf("Failed to AutoMigrate Table: %v, err: %v", table, err)
				return nil
			}
		case "article_infos":
			err = db.AutoMigrate(&model.ArticleInfo{})
			if err != nil {
				GetLogger().Errorf("Failed to AutoMigrate Table: %v, err: %v", table, err)
				return nil
			}
		}
	}
	return db
}
