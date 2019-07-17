package util

import (
	"github.com/gin-gonic/gin"
	"blog/pkg/e"
)

func Res(c *gin.Context, status int, code int) {
	c.JSON(status, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}