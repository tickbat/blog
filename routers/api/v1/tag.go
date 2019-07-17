package v1

import (
	"blog/models"
	"blog/pkg/e"
	"blog/pkg/setting"
	"blog/pkg/util"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}
	if arg := c.Query("state"); arg != "" {
		state := com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS
	data["lists"] = models.GetTags(util.GetPage(c), setting.Config.App.PageSize, maps)
	data["total"] = models.GetTagsTotal(maps)
	util.Res(c, http.StatusOK, code)
}

// 新增文章标签
func AddTag(c *gin.Context) {
	var tag models.Tag
	code := e.SUCCESS
	if err := c.ShouldBindJSON(&tag); err != nil {
		log.Printf("add tag parse json error: %v\n", err)
		code = e.INVALID_PARAMS
		util.Res(c, http.StatusBadRequest, code)
	} else {
		models.AddTag(tag)
		util.Res(c, http.StatusOK, code)
	}
}

// 修改文章标签
func EditTag(c *gin.Context) {
	var tag models.Tag
	var code = e.SUCCESS
	if err := c.ShouldBindJSON(&tag); err != nil {
		log.Printf("edit tag parse json error: %v\n", err)
		code := e.INVALID_PARAMS
		util.Res(c, http.StatusBadRequest, code)
		return
	}
	id := com.StrTo(c.Param("id")).MustInt()
	tag.ID = &id
	if models.ExistTagByID(id) {
		models.EditTag(tag)
		util.Res(c, http.StatusOK, code)
	} else {
		code = e.ERROR_NOT_EXIST_TAG
		util.Res(c, http.StatusOK, code)
	}
}

// 删除文章标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	code := e.SUCCESS
	if models.ExistTagByID(id) {
		models.DeleteTag(id)
		util.Res(c, http.StatusOK, code)
	} else {
		code = e.ERROR_NOT_EXIST_TAG
		util.Res(c, http.StatusOK, code)
	}
	util.Res(c, http.StatusOK, code)
}
