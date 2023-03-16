package typings

import (
	"sync"
	"time"
)

type SafeConversationMap struct {
	sync.RWMutex
	items map[string]*Conversation
}

func NewSafeConversationMap() *SafeConversationMap {
	return &SafeConversationMap{
		items: make(map[string]*Conversation),
	}
}

func (m *SafeConversationMap) Get(key string) *Conversation {
	m.RLock()
	defer m.RUnlock()
	return m.items[key]
}

func (m *SafeConversationMap) Set(key string, value *Conversation) {
	m.Lock()
	defer m.Unlock()
	m.items[key] = value
}

func (m *SafeConversationMap) Exists(key string) bool {
	m.RLock()
	defer m.RUnlock()
	_, ok := m.items[key]
	return ok
}

func (m *SafeConversationMap) Remove(key string) {
	m.Lock()
	defer m.Unlock()
	delete(m.items, key)
}

func (m *SafeConversationMap) All() map[string]*Conversation {
	m.RLock()
	defer m.RUnlock()
	items := make(map[string]*Conversation, len(m.items))
	for k, v := range m.items {
		items[k] = v
	}
	return items
}
func (m *SafeConversationMap) CleanUp() {
	m.Lock()
	defer m.Unlock()
	for k, v := range m.items {
		// Compare current time to last interaction time
		// If it's been more than 5 minutes, remove the conversation
		if time.Now().Unix()-v.LastInteraction > 300 {
			delete(m.items, k)
		}
	}
}
