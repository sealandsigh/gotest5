package main

import (
	"fmt"
	"time"
)

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err=", err)
			fmt.Println("发邮件给admin@163.com")
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println(res)
}

func main() {
	test()
	for {
		fmt.Println("main()下面的代码是...")
		time.Sleep(time.Second)
	}
}
