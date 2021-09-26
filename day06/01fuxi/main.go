package main

import "fmt"

func main() {
	var a interface{} // 定义一个空接口变量a

	a = 100
	// 如何判断a保存的值类型是什么?
	// 类型断言
	// x.(T)

	if v, ok := a.(int8); ok {
		fmt.Println("猜对了,a 是int8", v)
	} else {
		fmt.Println("猜错了不是int8")
	}

	// 2 switch
	switch v2 := a.(type) {
	case int8:
		fmt.Println("int8", v2)
	case int16:
		fmt.Println("int16", v2)
	case string:
		fmt.Println("string", v2)
	case int:
		fmt.Println("int", v2)
	default:
		fmt.Println("滚")
	}
}
