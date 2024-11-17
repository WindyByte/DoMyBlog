package main

import (
	"backend/core"
	"backend/global"
	"log"
)

func main() {
	// 初始化配置文件
	global.Config = core.InitConfig()
	global.DB = core.InitGorm()
	log.Printf("global.Config: %+v", global.Config)
	log.Printf("global.DB: %+v", global.DB)
}
