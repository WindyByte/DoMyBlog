package core

import (
	"backend/global"
	"backend/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	//var tables []string
	//result := db.Raw("SHOW TABLES").Scan(&tables)
	//if result.Error != nil {
	//	global.Log.Errorf("Failed to query tables: %v", result.Error)
	//}
	// 输出表列表
	//global.Log.Infof("Tables in database: %v", tables)
	//ClearAndCreateTable(db, "user_info", model.UserInfo{})
	//err = db.AutoMigrate(model.UserInfo{})
	//err = db.Migrator().CreateTable(model.UserInfo{})
	//if err != nil {
	//	return nil
	//}
	err = ClearAndCreateTable(db, model.UserInfo{})
	if err != nil {
		return nil
	}
	//a := db.Migrator().HasTable(&model.UserInfo{})
	//global.Log.Infof("HasTable: %v", a)
	//a = db.Migrator().HasTable(&model.ArticleInfo{})
	//global.Log.Warnf("HasTable: %v", a)
	//global.Log.Warnf("db.Migrator().CurrentDatabase(): %v", db.Migrator().CurrentDatabase())
	//if err != nil {
	//	return nil
	//}
	return db
}

func ClearAndCreateTable(db *gorm.DB, tableModel interface{}) error {
	// 检查表是否存在
	if !db.Migrator().HasTable(tableModel) {
		// 如果表不存在，则创建表
		err := db.Migrator().CreateTable(tableModel)
		if err != nil {
			global.Log.Errorf("Failed to create table: %v", err)
			return err
		}
		global.Log.Infof("Table UserInfo created successfully")
	} else {
		global.Log.Infof("Table UserInfo already exists")
	}
	return nil
}

func InitDBData(db *gorm.DB) {
	users := []*model.UserInfo{
		{UserID: "1000000000", UserName: "Alice", Passwd: "123456", Role: 1},
		{UserID: "1000000001", UserName: "Bob", Passwd: "123456", Role: 1},
		{UserID: "1000000002", UserName: "Charlie", Passwd: "123456", Role: 1},
		{UserID: "1000000003", UserName: "David", Passwd: "123456", Role: 1},
		{UserID: "1000000004", UserName: "Eve", Passwd: "123456", Role: 1},
		{UserID: "1000000005", UserName: "Frank", Passwd: "123456", Role: 1},
		{UserID: "1000000006", UserName: "Grace", Passwd: "123456", Role: 1},
		{UserID: "1000000007", UserName: "Helen", Passwd: "123456", Role: 1},
		{UserID: "1000000008", UserName: "Ivy", Passwd: "123456", Role: 1},
		{UserID: "1000000009", UserName: "Jack", Passwd: "123456", Role: 1},
		{UserID: "1000000010", UserName: "Kathy", Passwd: "123456", Role: 1},
		{UserID: "1000000011", UserName: "Liam", Passwd: "123456", Role: 1},
		{UserID: "1000000012", UserName: "Mia", Passwd: "123456", Role: 1},
		{UserID: "1000000013", UserName: "Noah", Passwd: "123456", Role: 1},
		{UserID: "1000000014", UserName: "Olivia", Passwd: "123456", Role: 1},
		{UserID: "1000000015", UserName: "Peter", Passwd: "123456", Role: 1},
		{UserID: "1000000016", UserName: "Quinn", Passwd: "123456", Role: 1},
		{UserID: "1000000017", UserName: "Rachel", Passwd: "123456", Role: 1},
		{UserID: "1000000018", UserName: "Sarah", Passwd: "123456", Role: 1},
		{UserID: "1000000019", UserName: "Tom", Passwd: "123456", Role: 1},
		{UserID: "1000000020", UserName: "Uma", Passwd: "123456", Role: 1},
		{UserID: "1000000021", UserName: "Vicky", Passwd: "123456", Role: 1},
		{UserID: "1000000022", UserName: "Wendy", Passwd: "123456", Role: 1},
		{UserID: "1000000023", UserName: "Xavier", Passwd: "123456", Role: 1},
		{UserID: "1000000024", UserName: "Yvonne", Passwd: "123456", Role: 1},
		{UserID: "1000000025", UserName: "Zoe", Passwd: "123456", Role: 1},
		{UserID: "1000000026", UserName: "Admin", Passwd: "123456", Role: 1},
	}
	err := db.Create(&users).Error
	if err != nil {
		global.Log.Errorf("Failed to create users: %v", err)
		return
	}
	global.Log.Infof("initDBData success")
}
