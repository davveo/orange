package controller

import (
	"net/http"
	"orange/model"

	"github.com/gin-gonic/gin"
)

func PointCategory(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	queryParams["cat_id"] = ctx.DefaultQuery("cat_id", "0")
	data := model.CreatePointCateGoryFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, data)
}
