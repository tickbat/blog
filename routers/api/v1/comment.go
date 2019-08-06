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

func GetComments(c *gin.Context) {
	var comment models.QueryComment
	if err := util.ValidateQuery(c, &comment); err != nil {
		return
	}
	data, err := service.GetComments(comment)
	if err != nil {
		logging.Error("get comment from service error:", err)
		util.Res(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	util.Res(c, http.StatusOK, e.SUCCESS, data)
}

func AddComment(c *gin.Context) {
	var comment models.Comment
	if err := util.ValidateJson(c, comment); err != nil {
		return
	}
	if !service.ExistArticleByID(comment.ArticleID) {
		util.Res(c, http.StatusBadRequest, e.ERROR_NOT_EXIST_ARTICLE, nil)
	}
	if err := service.AddComment(comment); err != nil {
		logging.Error("add comment from service error:", err)
		util.Res(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	util.Res(c, http.StatusOK, e.SUCCESS, nil)
}

func DeleteComment(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	if !service.ExistCommentByID(id) {
		util.Res(c, http.StatusBadRequest, e.ERROR_NOT_EXIST_COMMENT, nil)
	}
	if err := service.DeleteComment(id); err != nil {
		logging.Error("delete comment from service error:", err)
		util.Res(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	util.Res(c, http.StatusOK, e.SUCCESS, nil)
}
