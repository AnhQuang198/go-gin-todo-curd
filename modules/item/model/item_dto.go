package model

import (
	"errors"
	"social-todo-list/modules/item/entity"
	"social-todo-list/modules/item/enum"
)

var (
	ErrTitleIsBlank  = errors.New("title cannot be blank")
	ErrStatusIsBlank = errors.New("status cannot be blank")
	ErrItemDeleted   = errors.New("item is deleted")
)

type TodoItemRequest struct {
	Id          int              `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Status      *enum.ItemStatus `json:"status"`
}

type TodoItemResponse struct {
	entity.SQLModel
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Status      *enum.ItemStatus `json:"status"`
}

func (dto *TodoItemRequest) ToItemEntity() *entity.TodoItem {
	return &entity.TodoItem{
		Title:       dto.Title,
		Description: dto.Description,
		Status:      dto.Status,
	}
}

func FromEntity(item *entity.TodoItem) *TodoItemResponse {
	return &TodoItemResponse{
		Title:       item.Title,
		Description: item.Description,
		Status:      item.Status,
	}
}

func FromEntityList(items []*entity.TodoItem) []TodoItemResponse {
	var result []TodoItemResponse
	for _, item := range items {
		result = append(result, *FromEntity(item)) // dùng hàm FromEntity đã có
	}
	return result
}
