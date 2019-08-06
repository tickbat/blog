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

func GetReplies(c *gin.Context) {
	var reply models.QueryReply
	if err := util.ValidateQuery(c, &reply); err != nil {
		return
	}
	data, err := service.GetApplies(reply)
	if err != nil {
		util.Res(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	util.Res(c, http.StatusOK, e.SUCCESS, data)
}

func AddReply(c *gin.Context) {
	var reply models.Reply
	if err := util.ValidateQuery(c, &reply); err != nil {
		return
	}
	if !service.ExistCommentByID(reply.CommentID) {
		util.Res(c, http.StatusBadRequest, e.ERROR_NOT_EXIST_COMMENT, nil)
		return
	}
	if !service.ExistArticleByID(reply.ArticleID) {
		util.Res(c, http.StatusBadRequest, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}
	if err := service.AddReply(reply); err != nil {
		util.Res(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	util.Res(c, http.StatusOK, e.SUCCESS, nil)
}

func DeleteReply(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	if !service.ExistReplyByID(id) {
		util.Res(c, http.StatusBadRequest, e.ERROR_NOT_EXIST_REPLY, nil)
		return
	}
	if err := service.DeleteTag(id); err != nil {
		logging.Error("edit tag from service error:", err)
		util.Res(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	util.Res(c, http.StatusOK, e.SUCCESS, nil)
}
