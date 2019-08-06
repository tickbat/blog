package models

type Reply struct {
	Model
	CommentID   int    `json:"comment_id" binding:"required"`
	ArticleID   int    `json:"article_id" binding:"required"`
	Content     string `json:"article_id" binding:"required"`
	CreatedBy   string `json:"article_id" binding:"required"`
	Email       string `json:"article_id"`
	TargetName  string `json:"article_id" binding:"required"`
	TargetEmail string `json:"article_id" binding:"required"`
}

type QueryReply struct {
	CommentID int `json:"comment_id"`
}
