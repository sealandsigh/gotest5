package main

import (
	"fmt"
	"time"
)

// gorotine

func hello(i int) {
	fmt.Println("hello", i)
}

// 程序启动之后会创建一个主goroutine去执行
func main() {
	for i := 0; i < 100; i++ {
		// go hello(i) // 开启一个单独的goroutine 去执行hello函数(任务)
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	fmt.Println("main")
	time.Sleep(2 * time.Second)
	// main函数结束了 由main函数启动的goroutine 也都结束了
}
