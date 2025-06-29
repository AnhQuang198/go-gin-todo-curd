package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coder/websocket"
	"gorm.io/gorm"
	"log"
	"social-todo-list/modules/chatapp/business"
	"social-todo-list/modules/chatapp/model"
	"social-todo-list/modules/chatapp/storage"
	"time"
)

type Client struct {
	conn   *websocket.Conn
	roomId string
	userId string
	hub    *Hub
	send   chan string
}

// Tao moi 1 client
func NewClient(conn *websocket.Conn, roomId string, userId string, hub *Hub) *Client {
	return &Client{
		conn:   conn,
		roomId: roomId,
		userId: userId,
		send:   make(chan string, 256), //buffer 256 message
		hub:    hub,
	}
}

// Read nhận message từ client → broadcast
func (c *Client) Read(db *gorm.DB) {
	defer func() {
		c.hub.Leave(c.roomId, c)
		c.conn.Close(websocket.StatusNormalClosure, "bye")
	}()

	for {
		_, msg, err := c.conn.Read(context.Background())
		if err != nil {
			log.Println("read error:", err)
			break
		}

		var incoming model.MessageDTO
		if err := json.Unmarshal(msg, &incoming); err != nil {
			log.Println("invalid message:", err)
			continue
		}

		incoming.From = c.userId
		encoded, _ := json.Marshal(incoming)

		c.hub.Broadcast(c.roomId, string(encoded), c)

		//TODO: save message to DB
		store := storage.NewSQLStore(db)
		service := business.NewCreateMessageBusiness(store)
		service.CreateNewMessage(nil, c.roomId, c.userId, incoming.Content)
	}
}

// Write gửi message tới client
func (c *Client) Write() {
	for msg := range c.send {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		err := c.conn.Write(ctx, websocket.MessageText, []byte(msg))
		fmt.Println("Sending to:", c.userId, "msg:", msg)
		cancel()
		if err != nil {
			log.Println("write error:", err)
			break
		}
	}
}

// Send đưa message vào channel để gửi đi
func (c *Client) Send(msg string) {
	select {
	case c.send <- msg:
	default:
		c.conn.Close(websocket.StatusPolicyViolation, "slow client")
	}
}
