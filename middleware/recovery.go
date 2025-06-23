package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Recovery() func(*gin.Context) {
	return func(c *gin.Context) {
		log.Println("Recovery")
		c.Next()
	}
}
