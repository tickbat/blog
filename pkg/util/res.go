package util

import (
	"blog/pkg/e"
	"github.com/gin-gonic/gin"
)

func Res(c *gin.Context, status int, code int, data interface{}) {
	c.JSON(status, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
