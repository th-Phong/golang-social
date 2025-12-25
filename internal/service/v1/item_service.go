package v1service

import (
	"phongtran/go-social/golang-social/internal/db/sqlc"
	v1dto "phongtran/go-social/golang-social/internal/dto/v1"
	"phongtran/go-social/golang-social/internal/repository"

	"github.com/gin-gonic/gin"
)

type itemService struct {
	repo repository.ItemRepository
}

func NewItemService(repo repository.ItemRepository) ItemService {
	return &itemService{
		repo: repo,
	}
}

func (is *itemService) GetAllItems(ctx *gin.Context) ([]sqlc.TodoItem, error) {
	//context := ctx.Request.Context()

	//items, err := is.repo.GetAll(context)
	//if err != nil {
	//	return nil, err
	//}

	return []sqlc.TodoItem{}, nil
}

func (is *itemService) CreateItem(ctx *gin.Context, input v1dto.CreateItemRequest) (sqlc.TodoItem, error) {
	context := ctx.Request.Context()

	dbParams, err := input.MapCreateInputToParams()
	items, err := is.repo.Create(context, dbParams)
	if err != nil {
		return sqlc.TodoItem{}, err
	}

	return items, nil
}
