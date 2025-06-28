package model

import "sync"

type Hub struct {
	rooms map[string]map[*Client]bool //roomId - list client
	mutex sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		rooms: make(map[string]map[*Client]bool),
	}
}

// User join room
func (h *Hub) Join(roomId string, c *Client) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	if h.rooms[roomId] == nil {
		h.rooms[roomId] = make(map[*Client]bool)
	}
	h.rooms[roomId][c] = true
}

// User leave room
func (h *Hub) Leave(roomId string, c *Client) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	if clients, ok := h.rooms[roomId]; ok {
		delete(clients, c) //remove user khoi room
		if len(clients) == 0 {
			delete(h.rooms, roomId) //neu room khong con ai se remove room
		}
	}
}

// Broadcast: gửi tin nhắn tới tất cả client trong room, trừ sender
func (h *Hub) Broadcast(roomId string, msg string, sender *Client) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	for client := range h.rooms[roomId] {
		if client != sender {
			client.Send(msg)
		}
	}
}
