package service

import (
	"backend/core"
	"backend/infra"
	"backend/model"
	"backend/utils"
	"sort"
	"strings"
)

type ArticleInfo struct {
	ArticleID string `json:"articleID"`
	UserID    string `json:"userID"`
	Tags      string `json:"tags,omitempty"`    // 文章标签
	Title     string `json:"title"`             // 文章标题
	Content   string `json:"content,omitempty"` // 文章内容
	Views     string `json:"views"`             // 文章浏览量
	model.MetaInfo
}

func GetRelatedArticles(limit int) ([]ArticleInfo, error) {
	articleInfos := make([]ArticleInfo, 0)
	if limit <= 0 {
		return articleInfos, nil
	}
	articles, err := infra.GetArticles(limit)
	if err != nil {
		return articleInfos, err
	}
	for _, article := range articles {
		articleInfos = append(articleInfos, ArticleInfo{
			ArticleID: utils.Int64ToString(article.ArticleID),
			UserID:    utils.Int64ToString(article.UserID),
			Tags:      article.Tags,
			Title:     article.Title,
			Content:   article.Content,
			Views:     utils.Int64ToString(article.Views),
			MetaInfo:  article.MetaInfo,
		})
	}
	return articleInfos, nil
}

func GetArticleDetail(articleID int64) (ArticleInfo, error) {
	article, err := infra.GetArticleInfoByID(articleID)
	if err != nil {
		return ArticleInfo{}, err
	}
	return ArticleInfo{
		ArticleID: utils.Int64ToString(article.ArticleID),
		UserID:    utils.Int64ToString(article.UserID),
		Tags:      article.Tags,
		Title:     article.Title,
		Content:   article.Content,
		Views:     utils.Int64ToString(article.Views),
		MetaInfo:  article.MetaInfo,
	}, nil
}

func PublishArticle(title, content string, tagList []string) (int64, error) {
	articleID := utils.GenID()
	userID := utils.GenID()
	sort.Strings(tagList)
	tags := strings.Join(tagList, ",")
	articleInfo := model.ArticleInfo{
		ArticleID: articleID,
		UserID:    userID,
		Tags:      tags,
		Title:     title,
		Content:   content,
		Views:     0,
	}
	err := infra.CreateArticleInfo(articleInfo)
	if err != nil {
		core.GetLogger().Printf("CreateArticleInfo error, err: %v", err)
		return 0, err
	}
	return articleID, nil
}
