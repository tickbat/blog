package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"blog/models"
	"blog/pkg/e"
	"blog/pkg/util"
	"blog/pkg/setting"
	"net/http"
)

// 获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")
	var maps gin.H 
	var data gin.H
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

    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : data,
    })
}

// 新增文章标签
func AddTag(c *gin.Context) {

}

// 修改文章标签
func EditTag(c *gin.Context) {

}

// 删除文章标签
func DeleteTag(c *gin.Context) {

}