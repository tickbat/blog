package v1

import (
	"blog/models"
	"blog/pkg/e"
	"blog/pkg/logging"
	"blog/pkg/util"
	"blog/service"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTags(c *gin.Context) {
	var tag models.QueryTag
	if util.Validate(c, &tag) != nil {
		return
	}
	data, err := service.GetTags(&tag)
	if err != nil {
		logging.Error("get tags from service error:", err)
		util.Res(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	util.Res(c, http.StatusOK, e.SUCCESS, data)
}

// 新增文章标签
func AddTag(c *gin.Context) {
	var tag models.Tag
	if util.Validate(c, &tag) != nil {
		return
	}
	if err := service.AddTag(tag); err != nil {
		util.Res(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	util.Res(c, http.StatusOK, e.SUCCESS, nil)
}

// 修改文章标签
func EditTag(c *gin.Context) {
	var tag models.Tag
	if util.Validate(c, &tag) != nil {
		return
	}
	tag.ID = com.StrTo(c.Param("id")).MustInt()
	err := service.EditTag(tag)
	if err != nil {
		util.Res(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}
}

// 删除文章标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	exist, err := service.ExistTagByID(id)
	if err != nil {
		util.Res(c, http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}
	if !exist {
		util.Res(c, http.StatusBadRequest, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}
	if err := service.DeleteTag(id); err != nil {
		util.Res(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	util.Res(c, http.StatusOK, e.SUCCESS, nil)
}
