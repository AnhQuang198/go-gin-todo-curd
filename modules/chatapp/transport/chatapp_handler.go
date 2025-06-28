package transport

import (
	"github.com/coder/websocket"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"social-todo-list/modules/chatapp/model"
)

func JoinRoom() func(*gin.Context) {
	return func(c *gin.Context) {
		roomId := c.Query("room")
		userId := c.GetString("user_id")

		if roomId == "" || userId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing room or user"})
			return
		}

		conn, err := websocket.Accept(c.Writer, c.Request, &websocket.AcceptOptions{
			InsecureSkipVerify: true, // bỏ kiểm tra origin, nên cấu hình trong production
		})

		if err != nil {
			log.Println(err.Error())
			return
		}

		hub := model.NewHub()
		client := model.NewClient(conn, roomId, userId, hub)
		hub.Join(roomId, client)

		go client.Read()
		go client.Write()
	}
}
