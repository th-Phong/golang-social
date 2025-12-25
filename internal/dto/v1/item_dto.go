package v1dto

import "phongtran/go-social/golang-social/internal/db/sqlc"

type CreateItemParams struct {
	Title       string `json:"title" binding:"required,min=5"`
	Description string `json:"description,omitempty"`
	Image       []byte `json:"image,omitempty"`
	Status      int16  `json:"status,omitempty"`
}

type UpdateItemParams struct {
	Title       string  `json:"title" binding:"required,min=5"`
	Description *string `json:"description,omitempty"`
	Status      string  `json:"status"  binding:"required,oneof=Doing Done Deleted"`
}

type GetItemsParams struct {
	Search string `form:"search" binding:"omitempty,min=1,max=20"`
	Page   int    `form:"page" binding:"omitempty"`
	Limit  int    `form:"limit" binding:"omitempty"`
	Order  string `form:"order" binding:"omitempty"`
	Sort   string `form:"sort" binding:"omitempty"`
}

func (input *CreateItemParams) MapCreateInputToModel() sqlc.CreateItemParams {
	return sqlc.CreateItemParams{
		Title:       input.Title,
		Description: input.Description,
		Image:       input.Image,
		Status:      input.Status,
	}
}
