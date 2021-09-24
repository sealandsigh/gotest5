package main

import "fmt"

// 引出接口实例

// 定义一个能叫的类型
type speaker interface {
	speak() // 只要实现了speak方法的变量都是speaker这个类型,方法签名
}

type cat struct{}

type dog struct{}

type person struct{}

func (c cat) speak() {
	fmt.Println("喵喵喵")
}

func (d dog) speak() {
	fmt.Println("汪汪汪")
}

func (p person) speak() {
	fmt.Println("啊啊啊")
}

func da(x speaker) {
	// 接收一个参数，传递来干什么，我就打什么
	x.speak() // 挨打了就要叫
}

func main() {
	var c1 cat
	var d1 dog
	var p1 person

	da(c1)
	da(d1)
	da(p1)

}
