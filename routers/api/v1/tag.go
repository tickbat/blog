package v1

import (
	"blog/models"
	"blog/models/handler"
	"blog/pkg/e"
	"blog/pkg/logging"
	"blog/pkg/util"
	"blog/service"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTags(c *gin.Context) {
	tag := new(models.QueryTag)
	r := e.SUCCESS

	if err := c.ShouldBindQuery(tag); err != nil {
		logging.Info("get tags parse json error:", err)
		code := e.INVALID_PARAMS
		util.Res(c, http.StatusBadRequest, code, nil)
		return
	}

	data, err := service.GetTags(tag)
	if err != nil {
		logging.Info("tag_service.GetTags error:", err)
		r = e.ERROR
		util.Res(c, http.StatusInternalServerError, r, nil)
		return
	}
	util.Res(c, http.StatusOK, r, data)
}

// 新增文章标签
func AddTag(c *gin.Context) {
	var tag models.Tag
	code := e.SUCCESS
	if err := c.ShouldBindJSON(&tag); err != nil {
		logging.Info("add tag parse json error: " + err.Error())
		code = e.INVALID_PARAMS
		util.Res(c, http.StatusBadRequest, code, nil)
		return
	}
	if err := service.AddTag(tag); err != nil {
		code = e.ERROR
		util.Res(c, http.StatusInternalServerError, code, nil)
		return
	}
	util.Res(c, http.StatusOK, code, nil)
}

// 修改文章标签
func EditTag(c *gin.Context) {
	var tag models.Tag
	r := e.SUCCESS
	if err := c.ShouldBindJSON(&tag); err != nil {
		logging.Info("edit tag parse json error: " + err.Error())
		code := e.INVALID_PARAMS
		util.Res(c, http.StatusBadRequest, code, nil)
		return
	}
	id, err := com.StrTo(c.Param("id")).Int()
	if err != nil {
		r = e.INVALID_PARAMS
		util.Res(c, http.StatusBadRequest, r, nil)
		return
	}
	tag.ID = id
	err = service.EditTag(tag)
	if err == e.ERROR_NOT_EXIST_TAG {
		util.Res(c, http.StatusBadRequest, err, nil)
		return
	}
}

// 删除文章标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	r := e.SUCCESS
	if !models_handler.ExistTagByID(id) {
		r = e.ERROR_NOT_EXIST_TAG
		util.Res(c, http.StatusOK, r, nil)
		return
	}
	if err := service.DeleteTag(id); err != nil {
		r = e.ERROR
		util.Res(c, http.StatusInternalServerError, r, nil)
		return
	}
	util.Res(c, http.StatusOK, r, nil)
}
