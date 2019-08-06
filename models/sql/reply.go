package sql

import (
	"blog/models"
	"blog/pkg/logging"
	"github.com/jinzhu/gorm"
)

func GetReplies(reply models.QueryReply) ([]models.Reply, error) {
	var result []models.Reply
	if err := models.Db.Where(reply).Find(&result).Error; err != nil && err != gorm.ErrRecordNotFound {
		logging.Error("get reply from service error:", err)
		return result, err
	}
	return result, nil
}

func AddReply(reply models.Reply) error {
	err := models.Db.Create(reply).Error
	return err
}

func DeleteReply(id int) error {
	var reply models.Reply
	reply.ID = id
	err := models.Db.Delete(&reply).Error
	return err
}

func ExistReplyByID(id int) bool {
	var reply models.Reply
	if models.Db.First(&reply, "id = ?", id).RecordNotFound() {
		return false
	}
	return true
}
