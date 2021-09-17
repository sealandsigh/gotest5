package main

import "fmt"

// 函数: 一段代码的封装

func f5(title string, y ...int) int {
	fmt.Println(y) // y是一个int类型的切片
	return 1
}

func f6(x, y int) (sum int) {
	sum = x + y // 如果使用命名返回值，那么在函数中可以直接使用返回值变量
	return      // 如果使用命名返回值，return后面可以省略返回值变量
}

//go 语言中支持多个返回值
func f7(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func main() {
	f5("上海", 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(f6(1, 2))
	// 在一个命名函数中不能够再声明命名函数
	// func f8() {}
}
