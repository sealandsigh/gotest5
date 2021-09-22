package main

import "fmt"

// 自定义类型和类型别名

// type 后面跟的是类型

type myInt int     // 自定义类型
type yourInt = int // 类型别名

func main() {
	var n myInt = 100
	fmt.Println(n)
	var m yourInt = 100
	fmt.Println(m)

	var c rune = '中'
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
