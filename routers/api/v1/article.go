package v1

import (
	"github.com/gin-gonic/gin"
	"blog/models"
	"github.com/Unknwon/com"
	"blog/pkg/e"
)

//获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt
	code := e.SUCCESS
	models.GetArticle(id)
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