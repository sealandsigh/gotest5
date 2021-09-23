package main

import "fmt"

// 嵌套结构体

type address struct {
	province string
	city     string
}

// type workPlace struct {
// 	province string
// 	city     string
// }

type person struct {
	name string
	age  int
	// addr address
	address // 匿名嵌套结构体
	// workPlace //如果存在多个匿名嵌套结构体，并且字段有冲突，就需要完整的调用使用
}

type company struct {
	name string
	// address // 匿名嵌套结构体
	addr address
}

func main() {
	p1 := person{
		name: "周林",
		age:  9000,
		address: address{
			province: "山东",
			city:     "威海",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.address.city)
	fmt.Println(p1.city) // 先在自己结构体查找这个字段，找不到就去匿名嵌套的结构体中查找
}
