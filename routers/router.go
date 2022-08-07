package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mengdong123/go-gin-blog-jianyu/middleware/jwt"
	"github.com/mengdong123/go-gin-blog-jianyu/pkg/setting"
	"github.com/mengdong123/go-gin-blog-jianyu/pkg/upload"
	"github.com/mengdong123/go-gin-blog-jianyu/routers/api"
	"github.com/mengdong123/go-gin-blog-jianyu/routers/api/v1"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	r.GET("/auth", api.GetAuth)
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		// tag
		apiv1.GET("/tags", v1.GetTags)
		apiv1.GET("/tag/:id", v1.GetTag)
		apiv1.PUT("/tag/:id", v1.EditTag)
		apiv1.DELETE("/tag/:id", v1.DeleteTag)
		apiv1.POST("/tags", v1.AddTag)
		// article
		apiv1.POST("/article", v1.AddArticle)
		apiv1.DELETE("/article/:id", v1.DeleteArticle)
		apiv1.PUT("/article/:id", v1.EditArticle)
		apiv1.GET("/article/:id", v1.GetArticle)
		apiv1.GET("/articles", v1.GetArticles)

	}

	return r
}
