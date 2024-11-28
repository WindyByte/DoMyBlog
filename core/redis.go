package core

import (
	"github.com/redis/go-redis/v9"
)

func initRedisClient() *redis.Client {
	if GetBasicConfig().Redis.Addr == "" {
		GetLogger().Errorf("Redis addr is empty, skip init redis")
		return nil
	}
	// 初始化redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     GetBasicConfig().Redis.Addr,
		Password: GetBasicConfig().Redis.Password,
		DB:       GetBasicConfig().Redis.DB,
		PoolSize: GetBasicConfig().Redis.PoolSize,
	})
	return rdb
}
