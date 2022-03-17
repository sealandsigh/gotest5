package session

import (
	"errors"
	"sync"
)

// 对象
// MemorySession设计:
// 定义MemorySession 对象(字段 sessionId 存kv的map，读写锁)
// 构造函数，为了获取对象
// Set()
// Get()
// Del()
// Save()

type MemorySession struct {
	sessionId string
	data      map[string]interface{}
	rwlock    sync.RWMutex
}

// 构造函数

func NewMemorySession(id string) *MemorySession {
	s := &MemorySession{
		sessionId: id,
		data:      make(map[string]interface{}, 16),
	}
	return s
}

func (m *MemorySession) Set(key string, value interface{}) (err error) {
	// 加锁
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	// 设置值
	m.data[key] = value
	return
}

func (m *MemorySession) Get(key string, value interface{}) (err error) {
	// 加锁
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	// 设置值
	value, ok := m.data[key]
	if !ok {
		err = errors.New("key not exist in session")
		return
	}
	return
}

func (m *MemorySession) Del(key string, value interface{}) (err error) {
	// 加锁
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	delete(m.data, key)
	return
}

func (m *MemorySession) Save(key string, value interface{}) (err error) {
	return
}
