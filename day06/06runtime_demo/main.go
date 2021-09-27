package main

import (
	"fmt"
	"path"
	"runtime"
)

// runtime.Caller

func f() {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Printf("runtime.Caller failed\n")
	}
	// fmt.Println(pc)
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(file) //06runtime_demo/main.go
	fmt.Println(path.Base(file))
	fmt.Println(line)
}

func f1() {
	// _, file, line, ok := runtime.Caller(1)
	// if !ok {
	// 	fmt.Printf("runtime.Caller failed\n")
	// }
	// // fmt.Println(pc)
	// fmt.Println(file) //06runtime_demo/main.go
	// fmt.Println(line)
	f()
}

func main() {
	// _, file, line, ok := runtime.Caller(0)
	// if !ok {
	// 	fmt.Printf("runtime.Caller failed\n")
	// }
	// // fmt.Println(pc)
	// fmt.Println(file) //06runtime_demo/main.go
	// fmt.Println(line)
	f1()
}
