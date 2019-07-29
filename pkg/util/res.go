package util

import (
	"blog/pkg/e"
	"github.com/gin-gonic/gin"
)

func Res(c *gin.Context, status int, err error, data interface{}) {
	var code int
	r, ok := err.(e.Response)
	if ok {
		code = r.Code()
	} else {
		code = -1
	}
	c.JSON(status, gin.H{
		"code": code,
		"msg":  r.Error(),
		"data": data,
	})
}
