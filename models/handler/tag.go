package models_handler

import (
	"blog/models"
	"blog/models/cache"
	"blog/pkg/gredis"
	"encoding/json"
	"github.com/go-redis/redis"
	"blog/pkg/logging"
)

func GetTags(tag *models.QueryTag) ([]models.QueryTag, error) {
	var tagList []models.QueryTag
	key := cache_service.GetTagsKey(tag)
	val, err := gredis.Get(key)
	if err == redis.Nil {
		models.Db.Where(tag).Offset(tag.PageNum).Limit(tag.PageNum).Find(&tagList)
		gredis.Set(key, tagList, 3600)
		return tagList, nil
	}
	json.Unmarshal(val, &tagList)
	return tagList, err
}

func GetTagsTotal(maps interface{}) (count int) {
	models.Db.Model(&models.Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByID(id int) bool {
	var tag models.Tag
	models.Db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID == nil {
		return false
	}

	return true
}

func AddTag(tag models.Tag) error {
	if err := models.Db.Create(&tag).Error; err != nil {
		logging.Error("add tag from db error:", err)
		return err
	}
	return nil
}

func EditTag(tag models.Tag) {
	models.Db.Model(&tag).Update(tag)
	
}

func DeleteTag(id int) {
	tag := new(models.Tag)
	tag.ID = &id
	models.Db.Delete(tag)
}

func ClearAllTag() error {
	if err := models.Db.Where("deleted_at != null").Delete(&models.Tag{}).Error; err != nil {
		return err
	}
	return nil
}
