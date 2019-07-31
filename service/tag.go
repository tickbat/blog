package service

import (
	"blog/models"
	"blog/models/cache"
	"blog/models/sql"
	"blog/pkg/gredis"
	"blog/pkg/logging"
	"encoding/json"
)

func GetTags(tag *models.QueryTag) ([]models.QueryTag, error) {
	var tagList []models.QueryTag
	key := cache.GetTagsKey(tag)
	val, err := gredis.Get(key)
	if err != nil {
		logging.Warn("get tags from redis error:", err)
		return sql.GetTags(tag)
	}
	if err := json.Unmarshal(val, &tagList); err != nil {
		logging.Warn("get tags when parse json error:", err)
		return sql.GetTags(tag)
	}
	return tagList, nil
}

func AddTag(tag models.Tag) error {
	return sql.AddTag(tag)
}

func EditTag(tag models.Tag) error {
	return sql.EditTag(tag)
}

func DeleteTag(id int) error {
	return sql.DeleteTag(id)
}

func ExistTagByID(id int) (bool, error) {
	return sql.ExistTagByID(id)
}

func ClearAllTag() error {
	if err := models.Db.Where("deleted_at != null").Delete(&models.Tag{}).Error; err != nil {
		return err
	}
	return nil
}
