package main

import "fmt"

// 结构体遇到的问题
// 1 myInt(100)是什么？

type myInt int

func (m myInt) hello() {
	fmt.Println("我是一个Int")
}

func main() {
	m := myInt(100)
	m.hello()
	// 声明一个int32类型的变量x 它的值是10
	// 方法1
	// var x int32
	// x = 10
	// 方法2
	// var x int32 = 10
	// 方法3
	// var x = int32(10)
	// 方法4
	x := int32(10)
	fmt.Println(x)
}
