package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leryn1122/homepage/v2/pkg/config"
	"github.com/leryn1122/homepage/v2/pkg/middlewares"
	"github.com/leryn1122/homepage/v2/pkg/route"
)

func StartWebService() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.NoRoute(func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		method := ctx.Request.Method
		ctx.String(http.StatusNotFound, fmt.Sprintf("%s %s not found.", method, path))
	})

	router.Use(middlewares.Cors())
	route.ApiV1(router)

	_ = router.Run(fmt.Sprintf("%s:%d", config.Configuration.ServerConfig.Host, config.Configuration.ServerConfig.Port))
}
