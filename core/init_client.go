package core

import (
	"backend/config"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sync"
)

var (
	once        sync.Once
	gormClient  *gorm.DB
	log         *logrus.Logger
	redisClient *redis.Client
	basicConfig *config.BasicConfig
)

// GetBasicConfig 获取配置文件，单例模式
func GetBasicConfig() *config.BasicConfig {
	if basicConfig == nil {
		once.Do(initClient)
	}
	return basicConfig
}

// GetLogger 获取日志记录器
func GetLogger() *logrus.Logger {
	if log == nil {
		once.Do(initClient)
	}
	return log
}

// GetRedisClient 获取Redis客户端
func GetRedisClient() *redis.Client {
	if redisClient == nil {
		once.Do(initClient)
	}
	return redisClient
}
func GetGormClient() *gorm.DB {
	if gormClient == nil {
		once.Do(initClient)
	}
	return gormClient
}

func initClient() {
	basicConfig = initConfig()
	log = initLogger()
	redisClient = initRedisClient()
	gormClient = initGormClient()
}
