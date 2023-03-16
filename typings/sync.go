package typings

import "sync"

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
