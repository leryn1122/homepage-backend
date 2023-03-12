package route

import (
	"github.com/gin-gonic/gin"
	"github.com/leryn1122/homepage/v2/pkg/controller"
)

func ApiV1(router *gin.Engine) *gin.Engine {
	router.GET("/healthz", controller.Healthz)

	v1 := router.Group("/v1")
	{
		v1.GET("/sitemap", controller.ListSitemap)
	}
	return router
}
