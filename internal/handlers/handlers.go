package handlers

import (
	"github.com/acheong08/gpt4/api"
	"github.com/acheong08/gpt4/conversations"
	"github.com/acheong08/gpt4/typings"
	gin "github.com/gin-gonic/gin"
)

func NewConversation(c *gin.Context) {
	id := conversations.NewConversation()
	c.JSON(200, gin.H{
		"id": id,
	})
}

func AddEntry(c *gin.Context) {
	// Map the request body to the EntryRequest struct
	var entryRequest EntryRequest
	if err := c.BindJSON(&entryRequest); err != nil {
		c.JSON(400, gin.H{
			"error": "bad request",
		})
		return
	}
	if !conversations.RequestDataMap.Exists(entryRequest.ConversationID) {
		c.JSON(400, gin.H{
			"error": "conversation not found",
		})
		return
	}
	if entryRequest.Entry.Type == "text" {
		entryRequest.Entry.Data = entryRequest.Entry.Data + "<|im_end|>"
	}
	conversation := conversations.RequestDataMap.Get(entryRequest.ConversationID)
	conversation.RequestData.AddEntry(entryRequest.Entry)
	conversations.RequestDataMap.Set(entryRequest.ConversationID, conversation)
	c.JSON(200, gin.H{
		"success": true,
	})
}

func GetResponse(c *gin.Context) {
	conversationID := c.Param("conversation_id")
	if !conversations.RequestDataMap.Exists(conversationID) {
		c.JSON(400, gin.H{
			"error": "conversation not found",
		})
		return
	}
	conversation := conversations.RequestDataMap.Get(conversationID)
	response, err := api.Send(*conversation.RequestData)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "internal server error",
		})
		return
	}
	response.Model = ""
	c.JSON(200, response)
	c.Abort()

	var response_text string
	if len(response.Choices) > 0 {
		response_text = response.Choices[0].Text
	} else {
		response_text = "*Empty response*"
	}
	conversation.RequestData.AddEntry(typings.TranscriptEntry{
		Type: "text",
		Data: response_text + "<|im_end|>",
	})
	conversations.RequestDataMap.Set(conversationID, conversation)
}

func GetHistory(c *gin.Context) {
	conversationID := c.Param("conversation_id")
	if !conversations.RequestDataMap.Exists(conversationID) {
		c.JSON(400, gin.H{
			"error": "conversation not found",
		})
		return
	}
	conversation := conversations.RequestDataMap.Get(conversationID)
	c.JSON(200, conversation.RequestData.Transcript)
}
