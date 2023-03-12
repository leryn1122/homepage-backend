package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/leryn1122/homepage/v2/pkg/common"
	"net/http"
)

func Healthz(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &common.Result{
		Message: "Health check has passed!!",
	})
}
