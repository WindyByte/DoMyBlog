package main

import (
	"backend/core"
	"backend/global"
	"backend/routers"
)

//func CORSMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		// 设置允许跨域的头部信息
//		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // * 表示允许所有来源
//		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
//		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
//		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type")
//
//		// 如果是预检请求，直接返回
//		if c.Request.Method == "OPTIONS" {
//			c.AbortWithStatus(204)
//			return
//		}
//		c.Next()
//	}
//}

func main() {
	// 初始化配置文件
	global.Config = core.InitConfig()
	global.Log = core.InitLogger()
	global.RedisClient = core.InitRedisClient()
	global.MySqlClient = core.InitGormClient()
	core.InitDBData(global.MySqlClient)
	r := routers.InitRouter()

	global.Log.Info("Start backend server.")
	err := r.Run(global.Config.System.Addr())
	if err != nil {
		global.Log.Errorf("Start backend server failed, err: %v", err)
	}

}
