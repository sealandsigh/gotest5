package main

import (
	"fmt"
	"strings"
)

// 字符串

func main() {
	// \ 本来是具有特殊含义的, 我应该告诉程序我写的\是一个单纯的\
	path := "\"D:\\Go\\src\""
	fmt.Println(path)

	// 多行字符串
	s2 := `
		世情薄
				人情恶
		雨送黄昏花易落
	`
	fmt.Println(s2)

	// 字符串操作
	name := "理想"
	world := "dsb"

	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Printf("%s%s\n", name, world)
	fmt.Println(ss1)
	// 分隔
	ret := strings.Split(path, "\\")
	fmt.Println(ret)
	// 包含
	fmt.Println(strings.Contains(ss, "理性"))
	// 前缀
	fmt.Println(strings.HasPrefix(ss, "理想"))
	// 后缀
	fmt.Println(strings.HasSuffix(ss, "理想"))

	s4 := "abcdeb"
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.LastIndex(s4, "b"))
	// 拼接
	fmt.Println(strings.Join(ret, "+"))
}
