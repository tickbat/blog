package models

type Comment struct {
	Model
	Content   string `json:"content" binding:"required"`
	ArticleID int    `json:"articleId" binding:"required"`
	CreateBy  int    `json:"create_by" binding:"required"`
	Email     int    `json:"email"`
}

type QueryComment struct {
	ArticleID int `json:"article_id"`
}
