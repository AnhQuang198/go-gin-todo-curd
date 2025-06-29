package model

import "social-todo-list/modules/chatapp/entity"

// Message là struct gửi qua socket
type MessageDTO struct {
	From    string `json:"from"`
	Content string `json:"content"`
}

func FromEntity(item *entity.Message) *MessageDTO {
	return &MessageDTO{
		From:    item.SenderID,
		Content: item.Content,
	}
}

func FromEntityList(items []*entity.Message) []MessageDTO {
	var result []MessageDTO
	for _, item := range items {
		result = append(result, *FromEntity(item)) // dùng hàm FromEntity đã có
	}
	return result
}
