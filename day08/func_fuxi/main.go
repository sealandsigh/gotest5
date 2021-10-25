package main

import "fmt"

// 函数可变参数

func f1(a ...int) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func f2(a ...interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func main() {
	f1()
	f1(1)
	f1(1, 2, 3, 4, 5, 6, 7, 8, 9)
	f2(1, false, "a", []int{1, 2, 3}, [3]int{1, 2, 3}, map[string]int{"zhoulin": 9000}, 7, 8, 9)
}
