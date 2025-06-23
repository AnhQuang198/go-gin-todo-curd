package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"social-todo-list/middleware"
	ginitem "social-todo-list/modules/item/transport/gin"
)

func main() {
	dsn := os.Getenv("DN_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()

	//r.Use(middleware.Recovery()) //Ap dung cho toan bo Gin

	//CURD
	v1 := r.Group("/v1", middleware.Recovery()) //Ap dung cho 1 Group
	{
		items := v1.Group("/items")
		{
			items.POST("", ginitem.CreateItem(db)) //Ap dung cho tung API
			items.GET("", ginitem.ListItem(db))
			items.GET("/:id", ginitem.GetItem(db))
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8888")
}
