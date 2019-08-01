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
	if util.Validate(c, "json", &auth) != nil {
		return
	}
	ok, err := service.CheckAuth(auth)
	if err != nil {
		logging.Error("check auth from service error:", err)
		util.Res(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	if !ok {
		util.Res(c, http.StatusBadRequest, e.ERROR_AUTH, nil)
		return
	}
	token, err := util.GenerateToken(auth.Username)
	if err != nil {
		util.Res(c, http.StatusOK, e.ERROR_AUTH_TOKEN, err.Error())
		return
	}
	util.Res(c, http.StatusOK, e.SUCCESS, token)
}
