package main

import "fmt"

//

func main() {
	var a int = 100
	b := &a
	fmt.Printf("typea:%T typeb:%T\n", a, b)
	// 将a的十六进制内存地址打印出来
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", b)
	fmt.Printf("%v\n", b)
	fmt.Printf("%p\n", &b) // b的内存地址
}
