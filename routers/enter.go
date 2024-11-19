package routers

import (
	"backend/constant"
	"backend/global"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	r := gin.Default()
	loginRouter(r)
	articleRouter(r)
	return r
}

func pong(c *gin.Context) {
	c.JSON(200, constant.SuccessResp("pong"))
}
