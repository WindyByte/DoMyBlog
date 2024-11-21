package main

import (
	"backend/core"
	"backend/global"
	"backend/model"
)

func main() {
	// 初始化配置文件
	global.Config = core.InitConfig()
	global.Log = core.InitLogger()
	global.RedisClient = core.InitRedisClient()
	global.MySqlClient = core.InitGormClient()
	//r := routers.InitRouter()
	global.Log.Info("Start backend server.")
	//err := r.Run(global.Config.System.Addr())
	//core.InitDBData(global.MySqlClient)
	var users []model.UserInfo
	result := global.MySqlClient.Where("role = ?", 1).Find(&users)
	if result.Error != nil {
		global.Log.Errorf("Failed to query tables: %v", result.Error)
		return
	}
	global.Log.Infof("users: %v", users)
}
