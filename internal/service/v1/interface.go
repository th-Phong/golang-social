package v1service

import (
	"phongtran/go-social/golang-social/internal/db/sqlc"

	"github.com/gin-gonic/gin"
)

type ItemService interface {
	GetAllItems(ctx *gin.Context) ([]sqlc.TodoItem, error)
	CreateItem(ctx *gin.Context, input sqlc.CreateItemParams) (sqlc.TodoItem, error)
}
