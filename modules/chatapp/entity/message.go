package entity

import "time"

type Message struct {
	ID        int        `gorm:"primaryKey"`
	RoomID    string     `gorm:"column:room_id"`
	SenderID  string     `gorm:"column:sender_id"`
	Content   string     `gorm:"column:content"`
	CreatedAt *time.Time `gorm:"column:created_at"`
}

func (Message) TableName() string {
	return "messages"
}

func MappingData(roomId, senderId, message string) *Message {
	return &Message{
		RoomID:   roomId,
		SenderID: senderId,
		Content:  message,
	}
}
