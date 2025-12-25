package v1service

import (
	"errors"
	"phongtran/go-social/golang-social/internal/db/sqlc"
	v1dto "phongtran/go-social/golang-social/internal/dto/v1"
	"phongtran/go-social/golang-social/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
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

func (is *itemService) UpdateItem(ctx *gin.Context, input v1dto.UpdateItemRequest, id int32) (sqlc.TodoItem, error) {
	context := ctx.Request.Context()

	dbParams, err := input.MapUpdateInputToParams(id)

	items, err := is.repo.Update(context, dbParams)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return sqlc.TodoItem{}, errors.New("todo item not found or has been deleted")
		}
		return sqlc.TodoItem{}, err
	}

	return items, nil
}

func (is *itemService) GetItemDetail(ctx *gin.Context, id int32) (sqlc.TodoItem, error) {
	context := ctx.Request.Context()

	todoItem, err := is.repo.GetDetail(context, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return sqlc.TodoItem{}, errors.New("todo item not found or has been deleted")
		}
		return sqlc.TodoItem{}, err
	}
	return todoItem, nil
}

func (is *itemService) DeleteItem(ctx *gin.Context, id int32) error {
	cotext := ctx.Request.Context()

	err := is.repo.Delete(cotext, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.New("todo item not found or has been deleted")
		}
		return err
	}
	return nil
}

func (is *itemService) RestoreItem(ctx *gin.Context, id int32) (sqlc.TodoItem, error) {
	context := ctx.Request.Context()
	item, err := is.repo.Restore(context, id)

	if err != nil {
		return sqlc.TodoItem{}, err
	}
	return item, nil
}
