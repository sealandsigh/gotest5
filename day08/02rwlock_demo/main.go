package main

import (
	"fmt"
	"sync"
	"time"
)

// rwlock

var (
	x      = 0
	lock   sync.Mutex
	wg     sync.WaitGroup
	rwlock sync.RWMutex
)

// 读操作
func read() {
	defer wg.Done()
	rwlock.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
}

func write() {
	defer wg.Done()
	rwlock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 5)
	rwlock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		go write()
		wg.Add(1)
	}
	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	// 这里的时间差可以用time.Since替代就行
	// fmt.Println(time.Now().Sub(start))
	wg.Wait()
	fmt.Println(time.Since(start))
}
