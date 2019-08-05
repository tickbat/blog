package routers

import (
	"blog/middleware/jwt"
	"blog/pkg/setting"
	"blog/pkg/upload"
	"blog/routers/api"
	"blog/routers/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouters() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.App.RunMode)
	r.POST("/auth", api.GetAuth)
	r.POST("/upload", api.UploadImage)
	r.StaticFS("/upload/images", http.Dir(upload.GetImagePath()))

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tag", v1.AddTag)
		apiv1.PUT("/tag/:id", v1.EditTag)
		apiv1.DELETE("/tag/:id", v1.DeleteTag)
		apiv1.GET("/export", v1.ExportTag)
		apiv1.POST("/tags/import", v1.ImportTag)

		apiv1.GET("/article/:id", v1.GetArticle)
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.POST("/article", v1.AddArticle)
		apiv1.PUT("/article:id", v1.EditArticle)
		apiv1.DELETE("/article/:id", v1.DeleteArticle)

		apiv1.GET("/comments/:articleId", v1.GetComments)
		apiv1.POST("/comment", v1.AddComment)
		apiv1.DELETE("/comment/:id", v1.DeleteComment)
	}
	return r
}
