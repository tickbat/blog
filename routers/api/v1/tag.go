package v1

import (
	"log"
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

    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : data,
    })
}

// 新增文章标签
func AddTag(c *gin.Context) {
	var tag models.Tag
	if err := c.BindJSON(&tag); err != nil {
		log.Printf("%v\n", err)
		c.String(200, "失败了")
	} else {
		models.AddTag(tag)
		c.String(200, "成功了")
	}
}

// 修改文章标签
func EditTag(c *gin.Context) {

}

// 删除文章标签
func DeleteTag(c *gin.Context) {

}