package main

import (
	"backend/core"
	"backend/global"
	"backend/routers"
)

func main() {
	// 初始化配置文件
	global.Config = core.InitConfig()
	global.Log = core.InitLogger()
	global.RedisClient = core.InitRedisClient()
	global.MySqlClient = core.InitGormClient()
	r := routers.InitRouter()
	global.Log.Info("Start backend server.")
	err := r.Run(global.Config.System.Addr())
	if err != nil {
		return
	}

}
