package repository

import (
	"context"
	"phongtran/go-social/golang-social/internal/db/sqlc"
)

type ItemRepository interface {
	GetAll(ctx context.Context) ([]sqlc.TodoItem, error)
	Create(ctx context.Context, input sqlc.CreateItemParams) (sqlc.TodoItem, error)
	Update(ctx context.Context, input sqlc.UpdateItemParams) (sqlc.TodoItem, error)
}
