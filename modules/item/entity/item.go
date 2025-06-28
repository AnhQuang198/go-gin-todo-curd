package entity

import (
	"social-todo-list/modules/item/enum"
)

type TodoItem struct {
	SQLModel
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Status      *enum.ItemStatus `json:"status"`
}

func (TodoItem) TableName() string {
	return "todo_items"
}
