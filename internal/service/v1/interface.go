package v1service

import (
	"phongtran/go-social/golang-social/internal/db/sqlc"
	v1dto "phongtran/go-social/golang-social/internal/dto/v1"

	"github.com/gin-gonic/gin"
)

type ItemService interface {
	GetAllItems(ctx *gin.Context) ([]sqlc.TodoItem, error)
	CreateItem(ctx *gin.Context, input v1dto.CreateItemRequest) (sqlc.TodoItem, error)
	UpdateItem(ctx *gin.Context, input v1dto.UpdateItemRequest, id int32) (sqlc.TodoItem, error)
}
