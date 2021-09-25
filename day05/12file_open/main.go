package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 打开文件

// func readFromfile1() {
// 	fileObj, err := os.Open("./main.go")
// 	if err != nil {
// 		fmt.Printf("open file faild, err:%v", err)
// 		return
// 	}
// 	// 记得关闭文件
// 	defer fileObj.Close()
// 	// 读文件
// 	// var tmp := make([]byte, 128) // 指定读的长度
// 	var tmp [128]byte
// 	for {
// 		n, err := fileObj.Read(tmp[:])
// 		if err == io.EOF {
// 			fmt.Println("读完了")
// 			return
// 		}
// 		if err != nil {
// 			fmt.Printf("read file faild, err:%v", err)
// 			return
// 		}
// 		fmt.Printf("读了%d字节\n", n)
// 		fmt.Println(string(tmp[:n]))
// 		if n < 128 {
// 			return
// 		}
// 	}
// }

// func readFromFilebyBufio() {
// 	fileObj, err := os.Open("./main.go")
// 	if err != nil {
// 		fmt.Printf("open file faild, err:%v", err)
// 		return
// 	}
// 	// 记得关闭文件
// 	defer fileObj.Close()
// 	// 创建一个用来从文件中读内容的对象
// 	reader := bufio.NewReader(fileObj)
// 	for {
// 		line, err := reader.ReadString('\n')
// 		if err == io.EOF {
// 			return
// 		}
// 		if err != nil {
// 			fmt.Printf("read line faild,err:%v\n", err)
// 			return
// 		}
// 		fmt.Print(line)
// 	}
// }

func readFromFileByIoutil() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file faild, err:%v", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("read file faild,err:%v\n", err)
		return
	}
	fmt.Println(string(ret))
}

func main() {
	// readFromfile1()
	// readFromFilebyBufio()
	readFromFileByIoutil()
}
