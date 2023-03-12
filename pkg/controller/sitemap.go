package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/leryn1122/homepage/v2/pkg/common"
	"github.com/leryn1122/homepage/v2/pkg/database"
	"github.com/leryn1122/homepage/v2/pkg/model"
	"net/http"
)

func ListSitemap(ctx *gin.Context) {
	var sitemap []model.Sitemap
	err := database.DbClient.Limit(100).Where("enabled = ?", true).Find(&sitemap).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &common.Result{
			Code:    1,
			Message: "Internal server error",
		})
	}
	ctx.JSON(http.StatusOK, &common.Result{
		Data: sitemap,
	})
}
