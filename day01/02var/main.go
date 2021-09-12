package main

import "fmt"

// Go语言中推荐使用驼峰式的命名
// 声明变量
// var name string
// var age int
// var isok bool

// 批量声明
var (
	name string // ""
	age  int    // 0
	isok bool   // false
)

func main() {
	name = "理想"
	age = 23
	isok = true
	fmt.Print(isok)             // 在终端中输出要打印的内容，没有换行
	fmt.Printf("name:%s", name) // %s 是占位符 使用name这个变量的值替换占位符
	fmt.Println(age)            // 打印完指定的内容后会在后面加一个换行符

	// 声明变量同时赋值
	var s1 string = "whb"
	fmt.Print(s1)
	// 类型推导(根据值判断变量)
	var s2 = "20"
	fmt.Print(s2)
	// 简短变量声明,只能在函数里面使用
	s3 := "哈哈哈"
	fmt.Print(s3)
	// s1 := "10" 同一个作用域中不能声明重复的变量
}
