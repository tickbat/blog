package v1

import (
	"blog/models"
	"blog/pkg/e"
	"blog/pkg/setting"
	"blog/pkg/util"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"blog/pkg/logging"
	"net/http"
)

func GetTags(c *gin.Context) {
	data := make(map[string]interface{})
	tag := new(models.QueryTag)
	code := e.SUCCESS
	if err := c.ShouldBindQuery(tag); err != nil {
		logging.Info("get tags parse json error: " +  err.Error())
		code = e.INVALID_PARAMS
		util.Res(c, http.StatusBadRequest, code, nil)
		return
	}
	data["list"] = models.GetTags(util.GetPage(c), setting.App.PageSize, tag)
	data["total"] = models.GetTagsTotal(tag)
	util.Res(c, http.StatusOK, code, data)
}

// 新增文章标签
func AddTag(c *gin.Context) {
	var tag models.Tag
	code := e.SUCCESS
	if err := c.ShouldBindJSON(&tag); err != nil {
		logging.Info("add tag parse json error: " +  err.Error())
		code = e.INVALID_PARAMS
		util.Res(c, http.StatusBadRequest, code, nil)
	} else {
		models.AddTag(tag)
		util.Res(c, http.StatusOK, code, nil)
	}
}

// 修改文章标签
func EditTag(c *gin.Context) {
	var tag models.Tag
	var code = e.SUCCESS
	if err := c.ShouldBindJSON(&tag); err != nil {
		logging.Info("edit tag parse json error: " +  err.Error())
		code := e.INVALID_PARAMS
		util.Res(c, http.StatusBadRequest, code, nil)
		return
	}
	id := com.StrTo(c.Param("id")).MustInt()
	tag.ID = &id
	if models.ExistTagByID(id) {
		models.EditTag(tag)
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
	if !models.ExistTagByID(id) {
		code = e.ERROR_NOT_EXIST_TAG
		util.Res(c, http.StatusOK, code, nil)
		return
	}
	models.DeleteTag(id)
	util.Res(c, http.StatusOK, code, nil)
}
