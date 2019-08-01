package util

import (
	"blog/pkg/setting"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) (int, int) {
	page := 0
	size := setting.App.PageSize

	page = com.StrTo(c.Query("page")).MustInt()
	if page > 0 {
		page = (page - 1) * setting.App.PageSize
	}
	return page, size
}
