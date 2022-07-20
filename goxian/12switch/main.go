package main

import "fmt"

// 这是一个测试代码

func main() {
	var n1 int32 = 10
	var n2 int32 = 20
	switch n1 {
	case n2, 10, 5:
		fmt.Println("匹配成功!")
	default:
		fmt.Println("匹配失败!")
	}
}
