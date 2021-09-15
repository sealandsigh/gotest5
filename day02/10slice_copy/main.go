package main

import "fmt"

func main() {
	// a1 := []int{1, 3, 5}
	// a2 := a1
	// // var a3 []int 这里光var声明的话是copy不了的，因为是nil空 这里需要make
	// var a3 = make([]int, 3)
	// copy(a3, a1)
	// fmt.Println(a1, a2, a3)
	// a1[0] = 100
	// fmt.Println(a1, a2, a3)

	// // 将a1中的索引为1的3这个元素删除掉
	// a1 = append(a1[:1], a1[2:]...)
	// fmt.Println(a1)

	x1 := [...]int{1, 3, 5} // 数组
	s1 := x1[:]             // 切片
	fmt.Println(s1, len(s1), cap(s1))
	s1 = append(s1[:1], s1[2:]...)
	fmt.Println(s1, len(s1), cap(s1))
	// ?
	fmt.Println(x1)
}
