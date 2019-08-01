package service

import (
	"blog/models"
	"blog/models/cache"
	"blog/models/sql"
	"blog/pkg/gredis"
	"blog/pkg/logging"
	"encoding/json"
	"time"
)

func GetTags(tag *models.QueryTag, pageNum, pageSize int) ([]models.Tag, error) {
	var tagList []models.Tag
	key := cache.GetTagsKey(tag, pageNum, pageSize)
	val, err := gredis.Get(key)
	if err != nil && json.Unmarshal(val, &tagList) != nil {
		logging.Warn("get tags from redis error:", err)
		data, err := sql.GetTags(tag, pageNum, pageSize)
		if err != nil {
			return data, err
		}
		if err := gredis.Set(key, data, time.Second*60); err != nil {
			logging.Warn("set tags into redis error:", err)
		}
		return data, nil
	}
	return tagList, nil
	/*if err := json.Unmarshal(val, &tagList); err != nil {
		logging.Warn("get tags when parse json error:", err)
		data, err := sql.GetTags(tag, pageNum, pageSize)
		if err != nil {
			return data, err
		}
		if err := gredis.Set(key, data, time.Second * 60); err != nil {
			logging.Warn("set tags into redis error:", err)
		}
		return data, nil
	}*/
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

func GetTagsTotal(maps interface{}) (int, error) {
	return sql.GetTagsTotal(maps)
}

func ClearAllTag() error {
	if err := models.Db.Where("deleted_at != null").Delete(&models.Tag{}).Error; err != nil {
		return err
	}
	return nil
}
