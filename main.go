package main

import (
	"backend/core"
	"backend/global"
)

func main() {
	// 初始化配置文件
	global.Config = core.InitConfig()
	global.MySqlClient = core.InitGormClient()
	global.Log = core.InitLogger()
	global.RedisClient = core.InitRedisClient()
	global.Log.Info("Start backend server.")
}
