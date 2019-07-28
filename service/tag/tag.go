package tag_service

import (
	"blog/models"
	"blog/models/handler"
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
