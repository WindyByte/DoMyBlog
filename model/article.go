package model

type ArticleInfo struct {
	ArticleID int64  `gorm:"index;primaryKey;BIGINT;column:article_id" json:"articleID"`
	UserID    int64  `gorm:"index;BIGINT;column:user_id;not null" json:"userID"`
	Tags      string `gorm:"type:VARCHAR(255);column:tags" json:"tags,omitempty"`  // 文章标签
	Title     string `gorm:"type:VARCHAR(255);column:title;not null" json:"title"` // 文章标题
	Content   string `gorm:"type:TEXT;column:content" json:"content,omitempty"`    // 文章内容
	Views     int64  `gorm:"type:BIGINT;column:views;default:0" json:"views"`      // 文章浏览量
	MetaInfo
}
