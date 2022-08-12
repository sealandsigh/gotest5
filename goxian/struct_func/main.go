package main

import "fmt"

type Person struct {
	Name string
}

// 给Person结构体添加speak方法输出是一个好人
func (p Person) speak() {
	fmt.Println(p.Name, "是一个好人")
}
