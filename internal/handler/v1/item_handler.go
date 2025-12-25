package v1handler

import (
	"net/http"
	v1dto "phongtran/go-social/golang-social/internal/dto/v1"
	v1service "phongtran/go-social/golang-social/internal/service/v1"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	service v1service.ItemService
}

func NewItemHandler(service v1service.ItemService) *ItemHandler {
	return &ItemHandler{
		service: service,
	}
}

func (ih *ItemHandler) GetAllItem(ctx *gin.Context) {
	//var params v1dto.GetItemsParams
	//if err := ctx.ShouldBindQuery(&params); err != nil {
	//	return
	//}
	//
	//items, err := ih.service.GetAllItems(ctx)
	//
	//if err != nil {
	//	return
	//}
	//
	//ctx.JSON(http.StatusOK, gin.H{
	//	"data": items,
	//})
}

func (ih *ItemHandler) CreateItem(ctx *gin.Context) {
	var params v1dto.CreateItemRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	items, err := ih.service.CreateItem(ctx, params)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": v1dto.MapTodoResponse(items),
	})
}
