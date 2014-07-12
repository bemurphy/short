package short

import (
	"sync"
)

type Store interface {
	Set(string, string)
	Get(string) (string, bool)
	Del(string)
}

func NewMemoryStore() MemoryStore {
	return MemoryStore{data: make(map[string]string)}
}

type MemoryStore struct {
	data map[string]string
	m    sync.RWMutex
}

func (s MemoryStore) Set(key, val string) {
	s.m.Lock()
	s.data[key] = val
	s.m.Unlock()
}

func (s MemoryStore) Get(key string) (string, bool) {
	s.m.RLock()
	val, ok := s.data[key]
	s.m.RUnlock()

	return val, ok
}

func (s MemoryStore) Del(key string) {
	s.m.Lock()
	delete(s.data, key)
	s.m.Unlock()
}
