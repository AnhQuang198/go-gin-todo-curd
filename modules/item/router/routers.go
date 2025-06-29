package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"social-todo-list/middleware"
	"social-todo-list/modules/chatapp/transport"
	"social-todo-list/modules/chatapp/transport/websocket"
	ginitem "social-todo-list/modules/item/transport/gin"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	//r.Use(middleware.JWTAuthMiddleware()) //Ap dung authenJWT cho toan bo Gin

	//CURD
	v1 := r.Group("/v1", middleware.Recovery()) //Ap dung cho 1 Group
	{
		items := v1.Group("/items")
		{
			items.POST("", ginitem.CreateItem(db)) //Ap dung cho tung API
			items.GET("", ginitem.ListItem(db))
			items.GET("/:id", ginitem.GetItem(db))
			items.PATCH("/:id", ginitem.UpdateItem(db))
		}
	}

	hub := websocket.NewHub()
	chat := r.Group("/chatapp")
	{
		chat.GET("/ws", transport.JoinRoom(db, hub))
	}

	return r
}
