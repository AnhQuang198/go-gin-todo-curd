package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"social-todo-list/modules/item/router"
)

func main() {
	dsn := os.Getenv("DN_CONN_STR")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	r := router.SetupRouter(db)
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080" // fallback default port
	}
	_ = r.Run(port)
}
