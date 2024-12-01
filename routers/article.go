package routers

import (
	"backend/controller"
	"github.com/gin-gonic/gin"
)

func articleRouter(r *gin.Engine) {
	b := r.Group("/api/v1/article")
	{
		b.GET("/ping", pong)
		b.GET("/related", controller.GetRelateArticles)
		b.GET("/details", controller.GetArticleDetail)
		b.POST("/publish", controller.PublishArticle)
	}
}
