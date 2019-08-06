package v1

import (
	"blog/models"
	"blog/models/sql"
	"blog/pkg/e"
	"blog/pkg/logging"
	"blog/pkg/util"
	"blog/service"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	if !service.ExistArticleByID(id) {
		util.Res(c, http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	data, err := service.GetArticle(id)
	if err != nil {
		logging.Error("get article page from service error:", err)
		util.Res(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	util.Res(c, http.StatusOK, e.SUCCESS, data)
}

//获取多个文章
func GetArticles(c *gin.Context) {
	var article models.QueryArticle
	if util.ValidateQuery(c, &article) != nil {
		return
	}
	pageNum, pageSize := util.GetPage(c)
	list, err := service.GetArticles(&article, pageNum, pageSize)
	if err != nil {
		logging.Error("get articles from service error:", err)
		util.Res(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	page, err := service.GetTagsTotal(&article)
	if err != nil {
		logging.Error("get article page from service error:", err)
		util.Res(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	util.Res(c, http.StatusOK, e.SUCCESS, gin.H{
		"list":  list,
		"total": page,
	})
}

//新增文章
func AddArticle(c *gin.Context) {
	var article models.Article
	if util.ValidateJson(c, &article) != nil {
		return
	}
	if !service.ExistTagByID(article.TagID) {
		util.Res(c, http.StatusBadRequest, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}
	if err := service.AddArticle(article); err != nil {
		logging.Error("add article from service error:", err)
		util.Res(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	util.Res(c, http.StatusOK, e.SUCCESS, nil)
}

//修改文章
func EditArticle(c *gin.Context) {
	var article models.Article
	if util.ValidateJson(c, &article) != nil {
		return
	}
	article.ID = com.StrTo(c.Param("id")).MustInt()
	if !service.ExistArticleByID(article.ID) {
		util.Res(c, http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}
	if !service.ExistTagByID(article.TagID) {
		util.Res(c, http.StatusBadRequest, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}
	if err := sql.EditArticle(article); err != nil {
		logging.Error("edit article from service error:", err)
		util.Res(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	util.Res(c, http.StatusOK, e.SUCCESS, nil)
}

//删除文章
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	if !service.ExistArticleByID(id) {
		util.Res(c, http.StatusOK, e.ERROR, nil)
		return
	}
	if err := sql.DeleteArticle(id); err != nil {
		logging.Error("delete article from service error:", err)
		util.Res(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	util.Res(c, http.StatusOK, e.SUCCESS, nil)
}
