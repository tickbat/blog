package sql

import (
	"blog/models"
	"github.com/jinzhu/gorm"
)

func GetTags(tag *models.QueryTag, pageNum, pageSize int) ([]models.Tag, error) {
	var tagList []models.Tag
	if err := models.Db.Where(tag).Offset(pageNum).Limit(pageSize).Find(&tagList).Error; err != nil && err != gorm.ErrRecordNotFound {
		return tagList, err
	}
	return tagList, nil
}

func GetTagsAll(tag *models.QueryTag) ([]models.Tag, error) {
	var tagList []models.Tag
	if err := models.Db.Where(tag).Find(&tagList).Error; err != nil && err != gorm.ErrRecordNotFound {
		return tagList, err
	}
	return tagList, nil
}

func GetTagsTotal(maps interface{}) (int, error) {
	var count int
	err := models.Db.Model(&models.Tag{}).Where(maps).Count(&count).Error
	return count, err
}

func ExistTagByID(id int) bool {
	var tag models.Tag
	if models.Db.First(&tag, "id = ?", id).RecordNotFound() {
		return false
	}
	return true
}

func AddTag(tag models.Tag) error {
	err := models.Db.Create(&tag).Error
	return err
}

func EditTag(tag models.Tag) error {
	err := models.Db.Model(&tag).Update(tag).Error
	return err
}

func DeleteTag(id int) error {
	tag := new(models.Tag)
	tag.ID = id
	err := models.Db.Delete(tag).Error
	return err
}

/*func ClearAllTag() error {
	if err := models.Db.Where("deleted_at != null").Delete(&models.Tag{}).Error; err != nil {
		return err
	}
	return nil
}*/
