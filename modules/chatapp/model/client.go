package model

import (
	"context"
	"encoding/json"
	"github.com/coder/websocket"
	"log"
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
func (c *Client) Read() {
	defer func() {
		c.hub.Leave(c.roomId, c)
		c.conn.Close(websocket.StatusNormalClosure, "bye")
	}()

	ctx := context.Background()

	for {
		_, msg, err := c.conn.Read(ctx)
		if err != nil {
			log.Println("read error:", err)
			break
		}

		var incoming Message
		if err := json.Unmarshal(msg, &incoming); err != nil {
			continue
		}

		incoming.From = c.userId
		encoded, _ := json.Marshal(incoming)

		c.hub.Broadcast(c.roomId, string(encoded), c)
	}
}

// Write gửi message tới client
func (c *Client) Write() {
	for msg := range c.send {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		_ = c.conn.Write(ctx, websocket.MessageText, []byte(msg))
		cancel()
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
