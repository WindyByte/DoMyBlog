package routers

import "github.com/gin-gonic/gin"

func articleRouter(r *gin.Engine) {
	br := r.Group("/api/v1/blog")
	{
		br.GET("/ping", pong)
	}
}
