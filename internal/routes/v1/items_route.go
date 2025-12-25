package v1route

import (
	v1 "phongtran/go-social/golang-social/internal/handler/v1"

	"github.com/gin-gonic/gin"
)

type ItemRoute struct {
	handler *v1.ItemHandler
}

func NewItemRoute(handler *v1.ItemHandler) *ItemRoute {
	return &ItemRoute{
		handler: handler,
	}
}

func (ir ItemRoute) Register(r *gin.RouterGroup) {
	items := r.Group("/items")
	{
		items.GET("/", ir.handler.GetAllItem)
		items.GET("/:id", ir.handler.GetItemDetail)
		items.POST("/", ir.handler.CreateItem)
		items.PUT("/:id", ir.handler.UpdateItem)
		items.PUT("/:id/restore", ir.handler.RestoreItem)
		items.DELETE("/:id", ir.handler.DeleteItem)
	}
}
