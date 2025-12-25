package repository

import (
	"context"
	"log"
	"phongtran/go-social/golang-social/internal/db/sqlc"
)

type SQLItemRepository struct {
	db sqlc.Querier
}

func NewSQLItemRepository(db sqlc.Querier) ItemRepository {
	return &SQLItemRepository{
		db: db,
	}
}

func (ir *SQLItemRepository) GetAll(ctx context.Context) ([]sqlc.TodoItem, error) {
	return []sqlc.TodoItem{}, nil
}

func (ir *SQLItemRepository) Create(ctx context.Context, input sqlc.CreateItemParams) (sqlc.TodoItem, error) {
	todoItem, err := ir.db.CreateItem(ctx, input)
	if err != nil {
		return sqlc.TodoItem{}, err
	}
	return todoItem, nil
}

func (ir *SQLItemRepository) Update(ctx context.Context, input sqlc.UpdateItemParams) (sqlc.TodoItem, error) {
	todoItem, err := ir.db.UpdateItem(ctx, input)
	if err != nil {
		return sqlc.TodoItem{}, err
	}
	return todoItem, nil
}

func (ir *SQLItemRepository) GetDetail(ctx context.Context, id int32) (sqlc.TodoItem, error) {
	log.Printf("id: %d", id)
	todoItem, err := ir.db.GetItemDetail(ctx, id)
	if err != nil {
		return sqlc.TodoItem{}, err
	}
	return todoItem, nil
}

func (ir *SQLItemRepository) Delete(ctx context.Context, id int32) error {
	log.Printf("id: %d", id)
	_, err := ir.db.DeleteItem(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (ir *SQLItemRepository) Restore(ctx context.Context, id int32) (sqlc.TodoItem, error) {
	todoItem, err := ir.db.RestoreItem(ctx, id)
	if err != nil {
		return sqlc.TodoItem{}, err
	}
	return todoItem, nil
}
