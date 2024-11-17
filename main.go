package main

import (
	"backend/core"
	"backend/global"
	ctx "context"
)

func main() {
	// 初始化配置文件
	global.Config = core.InitConfig()
	global.MySqlClient = core.InitGormClient()
	global.Log = core.InitLogger()
	global.RedisClient = core.InitRedisClient()
	global.Log.Info("Start backend server.")
	//ok, err := global.RedisClient.Set(ctx.Background(), "test", "test_value", 10*time.Second).Result()
	//if err != nil {
	//	global.Log.Errorf("Set redis failed, err: %v", err)
	//	return
	//}
	//global.Log.Infof("Set redis success, ok: %v", ok)
	val, err := global.RedisClient.Get(ctx.Background(), "test").Result()
	if err != nil {
		global.Log.Errorf("Get redis failed: %v", err)
		return
	}
	global.Log.Infof("Get redis success, val: %v", val)
}
