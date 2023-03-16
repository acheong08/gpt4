package conversations

import (
	"github.com/acheong08/gpt4/typings"
	"github.com/google/uuid"
)

var (
	RequestDataMap = typings.NewSafeConversationMap()
)

func NewConversation() string {
	// Generate UUID
	id := uuid.New().String()
	RequestDataMap.Set(id, typings.Conversation{}.New())
	return id
}
