package sql

import "blog/models"

func GetComment(comment models.QueryComment) {
	var result models.Comment
	models.Db.Where(comment).Find(&result)
}

func delteComment(id int) error {
	var comment models.Comment
	comment.ID = id
	err := models.Db.Delete(&comment).Error
	return err
}
