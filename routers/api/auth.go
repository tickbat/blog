package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/validation"
	"blog/pkg/e"
	"blog/pkg/util"
	"blog/models"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	var auth models.Auth
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
	token, err := util.GenerateToken(username, password)
	if err != nil {
		code = e.ERROR_AUTH_TOKEN
		util.Res(c, http.SUCCESS, code, nil)
		return
	}
	util.Res(c, http.SUCCESS, code, token)
}