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
	exist, err := service.ExistArticleByID(id)
	if err != nil {
		util.Res(c, http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}
	if !exist {
		util.Res(c, http.StatusOK, e.ERROR_GET_ARTICLE_FAIL, nil)
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
	if util.Validate(c, "json", &article) != nil {
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
	if util.Validate(c, "json", &article) != nil {
		return
	}
	exist, err := service.ExistTagByID(article.TagId)
	if err != nil {
		logging.Error("test tag exist from service error:", err)
		util.Res(c, http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
	}
	if !exist {
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
	if util.Validate(c, "json", &article) != nil {
		return
	}
	article.ID = com.StrTo(c.Param("id")).MustInt()

	exist, err := service.ExistArticleByID(article.ID)
	if err != nil {
		util.Res(c, http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}
	if !exist {
		util.Res(c, http.StatusOK, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}

	exist, err = service.ExistTagByID(article.ID)
	if err != nil {
		logging.Error("test article exist from service error:", err)
		util.Res(c, http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exist {
		util.Res(c, http.StatusBadRequest, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}
	err = sql.EditArticle(article)
	if err != nil {
		logging.Error("edit article from service error:", err)
		util.Res(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
}

//删除文章
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	exist, err := service.ExistArticleByID(id)
	if err != nil {
		util.Res(c, http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}
	if !exist {
		util.Res(c, http.StatusOK, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}
	err = sql.DeleteArticle(id)
	if err != nil {
		logging.Error("delete article from service error:", err)
		util.Res(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
}
