package main

import "fmt"

// 函数

// 函数存在的意义？
// 函数是一段代码的封装
// 把一段逻辑抽象出来封装到一个函数中，给它起个名字，每次用到它的时候直接用函数名调用就可以了
// 使用函数能够让代码结构更清晰，更简洁

// 函数的定义
func sum(x int, y int) (ret int) {
	return x + y
}

// 没有返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

// 没有参数没有返回值
func f2() {
	fmt.Println("f2")
}

// 没有参数但有返回值
func f3() int {
	return 3
}

// 参数可以命名也可以不命名
// 命名的返回值就相当于在函数中声明一个变量
func f4(x int, y int) (ret int) {
	ret = x + y
	return
}

// 多个返回值
func f5() (int, string) {
	return 1, "沙河"
}

// 参数的类型简写, 当参数中连续两个参数的类型一致，我们可以将非最后一个参数的类型省略
func f6(x, y int) int {
	return x + y
}

// 可变长参数
// 可变长参数必须放在函数参数最后
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) // y的类型是切片
}

// Go语言中函数没有默认参数这个概念

// 参数可以命名也可以不命名
func main() {
	r := sum(1, 2)
	fmt.Println(r)

	_, n := f5()
	fmt.Println(n)

	f7("北京", 1)
	f7("上海", 1, 2, 3, 4, 5, 6, 7, 8, 9)

	f2()
	f1(1, 2)
	f3()
	f4(1, 2)
	f6(1, 2)

}
