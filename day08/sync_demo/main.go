package main

import (
	"fmt"
	"sync"
)

// 锁

// 这里用一个公共的空间可能就有问题，公共的x在两个goroutine变更的时候会有问题
// x变更有三步骤，从全局变量取x，x计算+1, x把值返回
var x = 0
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 50000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
