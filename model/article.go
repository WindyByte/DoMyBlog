package model

type ArticleInfo struct {
	MetaInfo
	Title   string   `json:"title"`
	Content string   `json:"content"`
	UserID  int64    `json:"user_id"`
	Tags    []string `json:"tags"`
}
