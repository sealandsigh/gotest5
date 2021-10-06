package main

import (
	"fmt"
	"sync"
)

// var a []int
var b chan int // 需要指定通道中元素的类型

var wg sync.WaitGroup

// func noBufChannel() {
// 	b = make(chan int) // 不带缓冲区的通道初始化
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		x := <-b
// 		fmt.Println("后台goroutine从管道b中取到了", x)
// 	}()
// 	b <- 10
// 	fmt.Println("10发送到了管道b中")
// 	wg.Wait()
// }

func BufChannel() {
	fmt.Println(b)
	b = make(chan int, 16) // 不带缓冲区的通道初始化
	b <- 10
	fmt.Println("10发送到了管道b中")
	close(b)
}

func main() {
	BufChannel()
}
