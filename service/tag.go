package service

import (
	"blog/models"
	"blog/models/handler"
	"blog/pkg/e"
)

func GetTags(tag *models.QueryTag) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	tagList, err := models_handler.GetTags(tag)
	if err != nil {
		return nil, err
	}
	page := models_handler.GetTagsTotal(tag)
	data["list"] = tagList
	data["total"] = page
	return data, err
}

func AddTag(tag models.Tag) error {
	return models_handler.AddTag(tag)
}

func EditTag(tag models.Tag) error {
	if !models_handler.ExistTagByID(tag.ID) {
		return e.ERROR_NOT_EXIST_TAG
	}
	return models_handler.EditTag(tag)
}

func DeleteTag(id int) error {
	return models_handler.DeleteTag(id)
}

func ClearAllTag() error {
	if err := models.Db.Where("deleted_at != null").Delete(&models.Tag{}).Error; err != nil {
		return err
	}
	return nil
}
