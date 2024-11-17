package main

import (
	"backend/core"
	"backend/global"
)

func main() {
	// 初始化配置文件
	global.Config = core.InitConfig()
	global.DB = core.InitGorm()
	global.Log = core.InitLogger()
	global.Log.Info("Start backend server ~")
}
