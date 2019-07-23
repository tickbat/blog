package v1

import (
	"blog/models"
	"blog/pkg/e"
	"blog/pkg/setting"
	"blog/pkg/util"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
	"blog/pkg/logging"
)

//获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	code := e.SUCCESS
	if !models.ExistArticleByID(id) {
		code = e.ERROR_NOT_EXIST_ARTICLE
		util.Res(c, http.StatusOK, code, nil)
		return
	}
	data := models.GetArticle(id)
	util.Res(c, http.StatusOK, code, data)
}

//获取多个文章
func GetArticles(c *gin.Context) {
	data := make(gin.H)
	var article models.QueryArticle
	code := e.SUCCESS
	if err := c.ShouldBindQuery(&article); err != nil {
		logging.Info("get articles parse json error: " +  err.Error())
		code = e.INVALID_PARAMS
		util.Res(c, http.StatusOK, code, nil)
		return
	}

	data["list"] = models.GetArticles(util.GetPage(c), setting.App.PageSize, article)
	data["total"] = models.GetArticlesTotal(article)
	util.Res(c, http.StatusOK, code, data)
}

//新增文章
func AddArticle(c *gin.Context) {
	var article models.Article
	code := e.SUCCESS
	if err := c.ShouldBindJSON(&article); err != nil {
		logging.Info("add article parse json error: " +  err.Error())
		code = e.INVALID_PARAMS
		util.Res(c, http.StatusOK, code, nil)
		return
	}
	if article.TagId != nil {
		if !models.ExistTagByID(*article.TagId) {
			code = e.ERROR_NOT_EXIST_TAG
			util.Res(c, http.StatusBadRequest, code, nil)
			return
		}
	}
	models.AddArticle(article)
	util.Res(c, http.StatusOK, code, nil)
}

//修改文章
func EditArticle(c *gin.Context) {
	var article models.Article
	code := e.SUCCESS
	if err := c.BindJSON(&article); err != nil {
		logging.Info("edit article parse json error: " +  err.Error())
		code = e.INVALID_PARAMS
		util.Res(c, http.StatusOK, code, nil)
		return
	}
	println(article.TagId)
	if article.TagId != nil {
		if models.ExistTagByID(*article.TagId) {
			code = e.INVALID_PARAMS
			util.Res(c, http.StatusBadRequest, code, nil)
			return
		}
	}
	models.EditArticle(article)
	util.Res(c, http.StatusOK, code, nil)
}

//删除文章
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	code := e.SUCCESS
	if !models.ExistArticleByID(id) {
		code = e.ERROR_NOT_EXIST_ARTICLE
		util.Res(c, http.StatusOK, code, nil)
		return
	}
	models.DeleteArticle(id)
	util.Res(c, http.StatusOK, code, nil)
}
