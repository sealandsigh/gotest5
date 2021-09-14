package main

import "fmt"

// switch
// 简化大量的判断
func main() {
	var n = 5
	if n == 1 {
		fmt.Println("大拇指")
	} else if n == 2 {
		fmt.Println("食指")
	} else if n == 5 {
		fmt.Println("小拇指")
	} else {
		fmt.Println("无效字符")
	}

	// switch 简化上面的代码
	switch n {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效字符")
	}
}
