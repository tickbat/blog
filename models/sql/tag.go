package sql

import (
	"blog/models"
	"blog/pkg/logging"
	"github.com/jinzhu/gorm"
)

func GetTags(tag *models.QueryTag) ([]models.QueryTag, error) {
	var tagList []models.QueryTag
	if err := models.Db.Where(tag).Offset(tag.PageNum).Limit(tag.PageNum).Find(&tagList).Error; err != nil && err != gorm.ErrRecordNotFound {
		return tagList, err
	}
	return tagList, nil
}

func GetTagsTotal(maps interface{}) (int, error) {
	var count int
	if err := models.Db.Model(&models.Tag{}).Where(maps).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

func ExistTagByID(id int) (bool, error) {
	var tag models.Tag
	if err := models.Db.Select("id").Where("id = ?", id).First(&tag).Error; err != nil {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

func AddTag(tag models.Tag) error {
	if err := models.Db.Create(&tag).Error; err != nil {
		logging.Error("add tag from db error:", err)
		return err
	}
	return nil
}

func EditTag(tag models.Tag) error {
	if err := models.Db.Model(&tag).Update(tag).Error; err != nil {
		logging.Error("edit tag from db error:", err)
		return err
	}
	return nil
}

func DeleteTag(id int) error {
	tag := new(models.Tag)
	tag.ID = id
	if err := models.Db.Delete(tag).Error; err != nil {
		logging.Error("delete tag from db error:", err)
		return err
	}
	return nil
}

/*func ClearAllTag() error {
	if err := models.Db.Where("deleted_at != null").Delete(&models.Tag{}).Error; err != nil {
		return err
	}
	return nil
}*/
