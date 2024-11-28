package routers

import (
	"backend/constant"
	"backend/core"
	"backend/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	gin.SetMode(core.GetBasicConfig().System.Env)
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	loginRouter(r)
	articleRouter(r)
	return r
}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, constant.SuccessResp("pong pong"))
}
