package routers

import (
	"backend/constant"
	"github.com/gin-gonic/gin"
)

func loginRouter(r *gin.Engine) {
	lr := r.Group("/api/v1/login")
	{
		lr.GET("/ping", login)
	}
}

func login(c *gin.Context) {
	c.JSON(200, constant.SuccessResp("login"))
}
