package main

import (
	"os"

	"github.com/acheong08/gpt4/internal/handlers"
	"github.com/gin-gonic/gin"
)

func authenticator(c *gin.Context) {
	if c.GetHeader("Authorization") != os.Getenv("AUTH_TOKEN") {
		c.AbortWithStatus(401)
		return
	}
	// Add CORS header allow all
	c.Header("Access-Control-Allow-Origin", "*")
	c.Next()
}

func main() {
	router := gin.Default()
	router.Use(authenticator)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// For all OPTIONS requests, return a 200
	router.OPTIONS("/*cors", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})
	router.GET("/conversation/new", handlers.NewConversation)
	router.POST("/conversation/add", handlers.AddEntry)
	router.GET("/conversation/:conversation_id/chat", handlers.GetResponse)
	router.GET("/conversation/:conversation_id/history", handlers.GetHistory)

	router.Run()

}
