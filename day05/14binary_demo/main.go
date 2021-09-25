package main

import "fmt"

// 二进制用途

const (
	eat   int = 4
	sleep int = 2
	da    int = 1
)

// 111
// 左边的1 吃饭 100
// 中间的1 睡觉 010
// 右边的1 打豆豆 001

func f(arg int) {
	fmt.Printf("%b\n", arg)
}

func main() {
	f(eat | da | sleep) // 101
}
