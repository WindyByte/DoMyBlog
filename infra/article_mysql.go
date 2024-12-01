package infra

import (
	"backend/core"
	"backend/model"
	"backend/utils"
)

func CreateArticleInfo(articleInfo model.ArticleInfo) error {
	db := core.GetGormClient()
	err := db.Create(&articleInfo).Error
	return err
}

func GetArticleInfoByID(ArticleID int64) (*model.ArticleInfo, error) {
	var articleInfo model.ArticleInfo
	db := core.GetGormClient()
	result := db.First(&articleInfo, ArticleID)
	return &articleInfo, result.Error
}

func GetArticles(limit int) ([]*model.ArticleInfo, error) {
	var articles []*model.ArticleInfo // 存储查询结果
	db := core.GetGormClient()        // 获取 GORM 客户端
	// 查询最新的 limit 条记录
	result := db.Limit(limit).Find(&articles)
	return articles, result.Error
}

func MInsertTestArticleInfo() error {
	db := core.GetGormClient()
	log := core.GetLogger()
	for i := 0; i < 10; i++ {
		article := model.ArticleInfo{
			ArticleID: utils.GenID(),
			Title:     "Article " + utils.IntToString(i),
			Content:   "Content for Article " + utils.IntToString(i),
			Tags:      "Tag" + utils.IntToString(i),
			UserID:    utils.GenID(),
		}
		if err := db.Create(&article).Error; err != nil {
			return err
		}
		log.Println("Test user info inserted successfully!")

	}
	return nil
}
