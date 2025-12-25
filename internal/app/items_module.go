package app

import (
	v1handler "phongtran/go-social/golang-social/internal/handler/v1"
	"phongtran/go-social/golang-social/internal/repository"
	"phongtran/go-social/golang-social/internal/routes"
	v1routes "phongtran/go-social/golang-social/internal/routes/v1"
	v1service "phongtran/go-social/golang-social/internal/service/v1"
)

type ItemModule struct {
	routes routes.Route
}

func NewItemModule(ctx *ModuleContext) *ItemModule {
	itemRepo := repository.NewSQLItemRepository(ctx.DB)

	itemService := v1service.NewItemService(itemRepo)

	itemHandler := v1handler.NewItemHandler(itemService)

	itemRoute := v1routes.NewItemRoute(itemHandler)

	return &ItemModule{
		routes: itemRoute,
	}
}

func (m *ItemModule) Routes() routes.Route {
	return m.routes
}
