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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	items, err := ih.service.CreateItem(ctx, params)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": v1dto.MapTodoResponse(items),
	})
}

func (ih *ItemHandler) UpdateItem(ctx *gin.Context) {
	var idParam v1dto.GetTodoIdParam
	if err := ctx.ShouldBindUri(&idParam); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var params v1dto.UpdateItemRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todoUpdate, err := ih.service.UpdateItem(ctx, params, idParam.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": v1dto.MapTodoResponse(todoUpdate),
	})
}

func (ih *ItemHandler) GetItemDetail(ctx *gin.Context) {
	var param v1dto.GetTodoIdParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	todoItem, err := ih.service.GetItemDetail(ctx, param.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": v1dto.MapTodoResponse(todoItem),
	})
}

func (ih *ItemHandler) DeleteItem(ctx *gin.Context) {
	var param v1dto.GetTodoIdParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ih.service.DeleteItem(ctx, param.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (ih *ItemHandler) RestoreItem(ctx *gin.Context) {
	var param v1dto.GetTodoIdParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todoItem, err := ih.service.RestoreItem(ctx, param.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": v1dto.MapTodoResponse(todoItem),
	})
}
