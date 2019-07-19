package routers

import (
	"blog/routers/api/v1"
	"blog/routers/api"
	"github.com/gin-gonic/gin"
	"blog/pkg/setting"
	"gin-blog/middleware/jwt"
)

func InitRouters() *gin.Engine {
	r := gin.New()
    r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.Config.RunMode)

	r.GET("/auth", api.GetAuth)
	
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tag", v1.AddTag)
		apiv1.PUT("/tag/:id", v1.EditTag)
		apiv1.DELETE("/tag/:id", v1.DeleteTag)

		apiv1.GET("/article/:id", v1.GetArticle)
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.POST("/article", v1.AddArticle)
		apiv1.PUT("/article", v1.EditArticle)
		apiv1.DELETE("/article", v1.DeleteArticle)
	}
	return r
}
