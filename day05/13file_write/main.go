package main

import (
	"fmt"
	"io/ioutil"
)

// 打开文件写内容

// func writeDemo1()  {
// 	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
// 	if err != nil {
// 		fmt.Printf("open file failed err:%v\n", err)
// 	}
// 	//write
// 	fileObj.Write([]byte("zhoulin mengbile!\n"))
// 	// writeString
// 	fileObj.WriteString("周林懵逼了")
// 	fileObj.Close()
// }

// func writeDemo2() {
// 	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
// 	if err != nil {
// 		fmt.Printf("open file failed err:%v\n", err)
// 	}
// 	defer fileObj.Close()
// 	// 创建一个写的对象
// 	wr := bufio.NewWriter(fileObj)
// 	wr.WriteString("hello沙河\n") // 写到缓存中
// 	wr.Flush()
// }

func writeDemo3() {
	str := "hello 沙河"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 6666)
	if err != nil {
		fmt.Printf("write file failed, err:%v\n", err)
		return
	}

}

func main() {
	// writeDemo1()
	// writeDemo2()
	writeDemo3()
}
