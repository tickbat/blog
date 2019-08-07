package api

import (
	"blog/models"
	"blog/pkg/e"
	"blog/pkg/logging"
	"blog/pkg/util"
	"blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAuth(c *gin.Context) {
	var auth models.Auth
	if util.ValidateJson(c, &auth) != nil {
		return
	}
	failType, err := service.CheckAuth(auth)
	if err != nil {
		logging.Error("check auth from service error:", err)
		util.Res(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	if failType == 1 {
		util.Res(c, http.StatusUnauthorized, e.ERROR_NOT_EXIST_USER, nil)
		return
	}
	if failType == 2 {
		util.Res(c, http.StatusUnauthorized, e.ERROR_MATCH_PASSWORD, nil)
		return
	}
	token, err := util.GenerateToken(auth.Username)
	if err != nil {
		logging.Error("generate token error:", err)
		util.Res(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	util.Res(c, http.StatusOK, e.SUCCESS, token)
}
