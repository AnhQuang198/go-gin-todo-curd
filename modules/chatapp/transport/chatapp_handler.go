package transport

import (
	"github.com/coder/websocket"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"social-todo-list/modules/chatapp/business"
	"social-todo-list/modules/chatapp/storage"
	websocket2 "social-todo-list/modules/chatapp/transport/websocket"
)

func JoinRoom(db *gorm.DB, hub *websocket2.Hub) func(*gin.Context) {
	return func(c *gin.Context) {
		roomId := c.Query("room")
		userId := c.Query("user_id")

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

		client := websocket2.NewClient(conn, roomId, userId, hub)
		hub.Join(roomId, client)

		store := storage.NewSQLStore(db)
		biz := business.NewGetMessageBusiness(store)
		data, _ := biz.GetMessage(nil)
		for i := len(data) - 1; i >= 0; i-- {
			client.Send(data[i].Content) // gửi tin theo đúng thứ tự
		}

		go client.Read(db)
		go client.Write()
	}
}
