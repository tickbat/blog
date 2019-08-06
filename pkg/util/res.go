package util

import (
	"blog/pkg/e"
	"blog/pkg/logging"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Res(c *gin.Context, status int, code int, data interface{}) {
	c.JSON(status, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func Validate(c *gin.Context, t string, obj interface{}) error {
	methods := map[string]func(obj interface{}) error{
		"json":  c.ShouldBindJSON,
		"query": c.ShouldBindQuery,
	}
	if err := methods[t](obj); err != nil {
		logging.Error("bind params error:", err)
		Res(c, http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return err
	}
	return nil
}

func ValidateJson(c *gin.Context, obj interface{}) error {
	return Validate(c, "json", obj)
}

func ValidateQuery(c *gin.Context, obj interface{}) error {
	return Validate(c, "query", obj)
}
