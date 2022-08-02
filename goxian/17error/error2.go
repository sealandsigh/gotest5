package main

import (
	"errors"
	"fmt"
)

func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误")
	}
}

func test02() {
	err := readConf("config3.ini")
	if err != nil {
		panic(err)
	}
	fmt.Println("test02继续执行")
}

func main() {
	test02()
	fmt.Println("main下面的代码是....")
}
