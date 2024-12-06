package main

import (
	"backend/core"
	"backend/routers"
)

func main() {
	// 初始化配置文件
	r := routers.InitRouter()
	core.GetLogger().Info("Start backend server.")
	err := r.Run(core.GetBasicConfig().System.Addr())
	if err != nil {
		core.GetLogger().Errorf("Start backend server failed, err: %v", err)
	}
}
