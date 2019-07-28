package v1

import (
	"blog/models"
	"blog/models/handler"
	"blog/pkg/e"
	"blog/pkg/logging"
	"blog/pkg/util"
	"blog/service/tag"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTags(c *gin.Context) {
	tag := new(models.QueryTag)
	code := e.SUCCESS

	if err := c.ShouldBindQuery(tag); err != nil {
		logging.Info("get tags parse json error:", err)
		code = e.INVALID_PARAMS
		util.Res(c, http.StatusBadRequest, code, nil)
		return
	}

	data, err := tag_service.GetTags(tag)
	if err != nil {
		logging.Info("tag_service.GetTags error:", err)
		code = e.ERROR
		util.Res(c, http.StatusInternalServerError, code, nil)
		return
	}
	util.Res(c, http.StatusOK, code, data)
}

// 新增文章标签
func AddTag(c *gin.Context) {
	var tag models.Tag
	code := e.SUCCESS
	if err := c.ShouldBindJSON(&tag); err != nil {
		logging.Info("add tag parse json error: " + err.Error())
		code = e.INVALID_PARAMS
		util.Res(c, http.StatusBadRequest, code, nil)
	} else {
		models_handler.AddTag(tag)
		util.Res(c, http.StatusOK, code, nil)
	}
}

// 修改文章标签
func EditTag(c *gin.Context) {
	var tag models.Tag
	var code = e.SUCCESS
	if err := c.ShouldBindJSON(&tag); err != nil {
		logging.Info("edit tag parse json error: " + err.Error())
		code := e.INVALID_PARAMS
		util.Res(c, http.StatusBadRequest, code, nil)
		return
	}
	id := com.StrTo(c.Param("id")).MustInt()
	tag.ID = &id
	if models_handler.ExistTagByID(id) {
		models_handler.EditTag(tag)
		util.Res(c, http.StatusOK, code, nil)
	} else {
		code = e.ERROR_NOT_EXIST_TAG
		util.Res(c, http.StatusOK, code, nil)
	}
}

// 删除文章标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	code := e.SUCCESS
	if !models_handler.ExistTagByID(id) {
		code = e.ERROR_NOT_EXIST_TAG
		util.Res(c, http.StatusOK, code, nil)
		return
	}
	models_handler.DeleteTag(id)
	util.Res(c, http.StatusOK, code, nil)
}
