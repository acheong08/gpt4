package handlers

import "github.com/acheong08/gpt4/typings"

type EntryRequest struct {
	ConversationID string                  `json:"conversation_id"`
	Entry          typings.TranscriptEntry `json:"entry"`
}
