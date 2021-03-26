package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/kalifun/go-programming-tour-book-blog/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", v1.Tag{}.Create)
		apiv1.DELETE("/tags/:id", v1.Tag{}.Delete)
		apiv1.PUT("/tags/:id", v1.Tag{}.Update)
		apiv1.PATCH("/tags/:id/state", v1.Tag{}.Update)
		apiv1.GET("/tags", v1.Tag{}.List)

		apiv1.POST("/articles", v1.Article{}.Create)
		apiv1.DELETE("/articles/:id", v1.Article{}.Delete)
		apiv1.PUT("/articles/:id", v1.Article{}.Update)
		apiv1.PATCH("/articles/:id/state", v1.Article{}.Update)
		apiv1.GET("/articles/:id", v1.Article{}.Get)
		apiv1.GET("/articles", v1.Article{}.List)
	}
	return r
}
