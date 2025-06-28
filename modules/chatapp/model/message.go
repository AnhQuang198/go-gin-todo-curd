package model

// Message là struct gửi qua socket
type Message struct {
	From    string `json:"from"`
	Content string `json:"content"`
}
