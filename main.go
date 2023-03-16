package main

import (
	"github.com/acheong08/gpt4/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/conversation/new", handlers.NewConversation)
	router.POST("/conversation/add", handlers.AddEntry)
	router.GET("/conversation/:conversation_id", handlers.GetResponse)

	router.Run()

}
