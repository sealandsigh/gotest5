package main

import "fmt"

// vscode 不支持 go module

//pointer

func main() {
	// 1 & 取地址
	n := 18
	p := &n
	fmt.Println(&n)
	fmt.Printf("%T\n", p) // *int表示的是int类型的指针，int类型的内存地址

	// 2 * 根据地址取值
	m := *p
	fmt.Println(m)

	var a1 *int
	fmt.Println(a1)
	var a2 = new(int) //new函数申请新的内存地址
	fmt.Println(a2)
	*a2 = 100
	fmt.Println(*a2)

}
