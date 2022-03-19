package session

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"sync"
)

// MemorySessionMgr 设计
// 定义MemorySessionMgr对象(字段，存放所有session的map 读写锁)
// 构造函数
// Init()
// CreateSession()
// GetSession()

// 定义对象

type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwlock     sync.RWMutex
}

// 构造函数

func NewMemorySessionMgr() *MemorySessionMgr {
	sr := &MemorySessionMgr{
		sessionMap: make(map[string]Session),
	}
	return sr
}

func (s *MemorySessionMgr) Init(addr string, option ...string) (err error) {
	return
}

func (s *MemorySessionMgr) CreateSession() (session Session, err error) {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()
	// go get github.com/satori/go.uuid
	// 用uuid作为sessionId
	//id,err := uuid.NewV4()
	id := uuid.NewV4()
	if err != nil {
		return
	}
	// 转string
	sessionId := id.String()
	// 创建个session
	session = NewMemorySession(sessionId)
	// 加入到大的map中
	s.sessionMap[sessionId] = session
	return
}

func (s *MemorySessionMgr) Get(sessionId string) (session Session, err error) {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()
	session, ok := s.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not exists")
		return
	}
	return
}
