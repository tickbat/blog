package models

type Comment struct {
	model
	Content   string `json:"content" binding:"required"`
	ArticleId int    `json:"articleId" binding:"required"`
	CreateBy  int    `json:"create_by" binding:"required"`
	Email     int    `json:"email"`
}

type QueryComment struct {
	ArticleId int `json:"articleId" binding:"required"`
}
