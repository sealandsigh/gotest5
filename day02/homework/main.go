package main

import "fmt"

// 单行注释
/* 多行
注释
*/

// main函数是入口函数
// 没有参数也没有返回值
func main() {
	//打印 9x9乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, j*i)
		}
		fmt.Println()
	}
}
