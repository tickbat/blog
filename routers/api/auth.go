package api

import (
	"blog/models"
	"blog/pkg/e"
	"blog/pkg/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetAuth(c *gin.Context) {
	var auth models.Auth
	code := e.SUCCESS
	if err := c.ShouldBindJSON(&auth); err != nil {
		log.Printf("query auth parse json error: %v\n", err)
		code = e.INVALID_PARAMS
		util.Res(c, http.StatusBadRequest, code, nil)
		return
	}
	if !models.CheckAuth(auth) {
		code = e.ERROR_AUTH
		util.Res(c, http.StatusBadRequest, code, nil)
		return
	}
	token, err := util.GenerateToken(*auth.Username)
	if err != nil {
		code = e.ERROR_AUTH_TOKEN
		util.Res(c, http.StatusOK, code, err.Error())
		return
	}
	util.Res(c, http.StatusOK, code, token)
}
