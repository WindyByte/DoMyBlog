package global

import (
	"backend/config"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config      *config.Config
	MySqlClient *gorm.DB
	Log         *logrus.Logger
	RedisClient *redis.Client
)
