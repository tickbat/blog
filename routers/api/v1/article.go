package v1

import (
	"blog/models"
	"blog/pkg/e"
	"blog/pkg/util"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	code := e.SUCCESS
	data := models.GetArticle(id)
	util.Res(c, http.StatusOK, code, data)
}

//获取多个文章
func GetArticles(c *gin.Context) {
	code := e.SUCCESS
	data := models.GetArticles()

}

//新增文章
func AddArticle(c *gin.Context) {

}

//修改文章
func EditArticle(c *gin.Context) {

}

//删除文章
func DeleteArticle(c *gin.Context) {

}
