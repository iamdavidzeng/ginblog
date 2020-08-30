package routers

import (
	"gin-blog/middleware/jwt"
	"gin-blog/pkg/setting"
	"gin-blog/routers/api"
	v1 "gin-blog/routers/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article map[string]string

// InitRouter init router for server
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)

	article := Article{
		"title":   "1st title",
		"content": "1st content",
		"author":  "david",
	}

	r.LoadHTMLGlob("templates/**/*")
	r.GET("/articles/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "articles/index.tmpl", gin.H{
			"articles": []Article{article, article},
		})
	})

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
