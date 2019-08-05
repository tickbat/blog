package service

import (
	"blog/models"
	"blog/models/sql"
)

func GetComments(comment models.QueryComment) ([]models.Comment, error) {
	return sql.GetComment(comment)
}

func AddComment(comment models.Comment) error {
	return sql.AddComment(comment)
}

func DeleteComment(id int) error {
	return sql.DeleteComment(id)
}
