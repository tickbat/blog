package sql

import (
	"blog/models"
	"github.com/jinzhu/gorm"
)

func GetComment(comment models.QueryComment) ([]models.Comment, error) {
	var result []models.Comment
	if err := models.Db.Where(comment).Find(&result).Error; err != nil && err != gorm.ErrRecordNotFound {
		return result, err
	}
	return result, nil
}

func AddComment(comment models.Comment) error {
	err := models.Db.Create(comment).Error
	return err
}

func DeleteComment(id int) error {
	var comment models.Comment
	comment.ID = id
	err := models.Db.Delete(&comment).Error
	return err
}
