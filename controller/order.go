package controller

import (
	"net/http"
	"orange/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func OrderList(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	keyWord := ctx.Query("keyword")
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["page_no"] = pageNo
	queryParams["keyword"] = keyWord
	queryParams["page_size"] = pageSize
	data, dataTotal := model.CreateOrderFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}
