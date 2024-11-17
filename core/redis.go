package core

import (
	"backend/global"
	"github.com/redis/go-redis/v9"
)

func InitRedisClient() *redis.Client {
	if global.Config.Redis.Addr == "" {
		global.Log.Errorf("Redis addr is empty, skip init redis")
		return nil
	}
	// 初始化redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Addr,
		Password: global.Config.Redis.Password,
		DB:       global.Config.Redis.DB,
		PoolSize: global.Config.Redis.PoolSize,
	})
	return rdb
}
