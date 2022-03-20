package session

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

type RedisSessionMgr struct {
	// redis地址
	addr string
	// 密码
	password string
	// 连接池
	pool *redis.Pool
	// 锁
	rwlock sync.RWMutex
	// 大map
	sessionMap map[string]Session
}

func (r *RedisSessionMgr) Init(addr string, option ...string) (err error) {
	// 若有其它参数
	if len(option) > 0 {
		r.password = option[0]
	}
	// 创建连接池
	r.pool = myPool(addr, r.password)
	r.addr = addr
	return
}

func myPool(addr, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     64,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			// 若有密码,判断
			if _, err := conn.Do("AUTH", password); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, err
		},
		// 连接测试，开发时写
		// 上线注释掉
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}
}

func (r *RedisSessionMgr) CreateSession() (session Session, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
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
	session = NewRedisSession(sessionId, r.pool)
	// 加入到大的map中
	r.sessionMap[sessionId] = session
	return
}

func (r *RedisSessionMgr) Get(sessionId string) (session Session, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	session, ok := r.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not exists")
		return
	}
	return
}

// 构造函数
func NewRedisSessionMgr() SessionMgr {
	sr := &RedisSessionMgr{
		sessionMap: make(map[string]Session, 32),
	}
	return sr
}
