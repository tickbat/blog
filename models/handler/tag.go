package models_handler

import (
	"blog/models"
	"blog/models/cache"
	"blog/pkg/gredis"
	"encoding/json"
	"github.com/go-redis/redis"
)

func GetTags(tag *models.QueryTag) ([]models.QueryTag, error) {
	var resList []models.QueryTag
	key := cache_service.GetTagsKey(tag)
	val, err := gredis.Get(key)
	if err == redis.Nil {
		models.Db.Where(tag).Offset(tag.PageNum).Limit(tag.PageNum).Find(&resList)
		return resList, nil
	}
	json.Unmarshal(val, &resList)
	return resList, err
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

func AddTag(tag models.Tag) bool {
	models.Db.Create(&tag)
	return true
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
