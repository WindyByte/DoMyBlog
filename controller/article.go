package controller

import (
	"backend/constant"
	"backend/service"
	"backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRelateArticles(c *gin.Context) {
	limitQuery := c.DefaultQuery("limit", "5")
	limit, err := utils.StringToInt(limitQuery)
	if err != nil {
		c.JSON(400, constant.FailResp(fmt.Sprintf("StringToInt error, err: %v", err)))
		return
	}
	articles, err := service.GetRelatedArticles(limit)
	if err != nil {
		c.JSON(http.StatusForbidden, constant.FailResp(fmt.Sprintf("GetArticles error, err: %v", err)))
		return
	}
	c.JSON(http.StatusOK, constant.SuccessResp(articles))
}

func GetArticleDetail(c *gin.Context) {
	articleIDQuery := c.DefaultQuery("articleID", "-1")
	articleID, err := utils.StringToInt64(articleIDQuery)
	if err != nil {
		c.JSON(400, constant.FailResp(fmt.Sprintf("StringToInt64 error, err: %v", err)))
		return
	}
	article, err := service.GetArticleDetail(articleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, constant.FailResp(fmt.Sprintf("GetArticleInfoByID error, err: %v", err)))
		return
	}
	c.JSON(http.StatusOK, constant.SuccessResp(article))
}

type PublishArticleReq struct {
	Title   string   `json:"title"`          // 文章标题
	Html    string   `json:"html"`           // HTML渲染后的文章内容
	Content string   `json:"markdown"`       // 文章内容
	Tags    []string `json:"tags,omitempty"` // 文章标签
}

func PublishArticle(c *gin.Context) {
	var req PublishArticleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, constant.FailResp(fmt.Sprintf("ShouldBindJSON error, err: %v", err)))
		return
	}
	articleID, err := service.PublishArticle(req.Title, req.Content, req.Tags)
	if err != nil {
		c.JSON(http.StatusBadRequest, constant.FailResp(fmt.Sprintf("PublishArticle error, err: %v", err)))
		return
	}
	c.JSON(http.StatusOK, constant.SuccessResp(articleID))
}
