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
		"msg":  e.GetMsg,
		"data": data,
	})
}

func Validate(c *gin.Context, params interface{}) error {
	if err := c.ShouldBindJSON(params); err != nil {
		logging.Info("bind params error:", err)
		Res(c, http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return err
	}
	return nil
}
