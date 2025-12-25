package v1dto

import (
	"encoding/json"
	"phongtran/go-social/golang-social/internal/db/sqlc"
	"time"
)

type GetTodoIdParam struct {
	ID int32 `uri:"id" binding:"min=1,max=2147483647"`
}

type ImageMetadata struct {
	URL      string `json:"url"`
	FileName string `json:"file_name"`
	Size     *int64 `json:"size"`
}

type CreateItemRequest struct {
	Title       string         `json:"title" binding:"required,min=5"`
	Description string         `json:"description,omitempty"`
	Image       *ImageMetadata `json:"image,omitempty"`
	Status      int16          `json:"status,omitempty" binding:"omitempty,oneof=1 2 3"`
}

type UpdateItemRequest struct {
	Title       *string        `json:"title" binding:"omitempty,required,min=5"`
	Description string         `json:"description,omitempty"`
	Image       *ImageMetadata `json:"image,omitempty"`
	Status      *int16         `json:"status" binding:"omitempty,oneof=1 2 3"`
}

type ItemResponse struct {
	ID          int32          `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Image       *ImageMetadata `json:"image"`
	Status      int16          `json:"status"`
	CreatedAt   string         `json:"created_at"`
	UpdatedAt   string         `json:"updated_at"`
}

func (req *CreateItemRequest) MapCreateInputToParams() (sqlc.CreateItemParams, error) {
	var imgBytes []byte = nil
	var err error

	if req.Image != nil && req.Image.URL != "" {
		imgBytes, err = json.Marshal(req.Image)
		if err != nil {
			return sqlc.CreateItemParams{}, err
		}
	}

	return sqlc.CreateItemParams{
		Title:       req.Title,
		Description: req.Description,
		Image:       imgBytes,
		Status:      req.Status,
	}, nil
}

func (req *UpdateItemRequest) MapUpdateInputToParams(id int32) (sqlc.UpdateItemParams, error) {
	var imgBytes []byte = nil
	var err error

	if req.Image != nil && req.Image.URL != "" {
		imgBytes, err = json.Marshal(req.Image)
		if err != nil {
			return sqlc.UpdateItemParams{}, err
		}
	}

	return sqlc.UpdateItemParams{
		Title:       req.Title,
		Description: req.Description,
		Image:       imgBytes,
		Status:      req.Status,
		ID:          id,
	}, nil
}

func MapTodoResponse(todo sqlc.TodoItem) *ItemResponse {
	var img *ImageMetadata = nil

	if len(todo.Image) > 0 {
		_ = json.Unmarshal(todo.Image, &img)
	}

	return &ItemResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Image:       img,
		Status:      todo.Status,
		CreatedAt:   todo.CreatedAt.Time.Format(time.DateTime),
		UpdatedAt:   todo.UpdatedAt.Time.Format(time.DateTime),
	}
}
