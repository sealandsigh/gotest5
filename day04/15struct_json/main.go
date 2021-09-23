package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与json

// 1 把go语言中的结构体变量----》 json格式的字符串  序列化
// 2 json格式的字符串 ---》go语言中能够识别的结构体变量 反序列化

type person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "周林",
		Age:  18,
	}
	fmt.Println(p1)
	// 序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("Marshal fauld err:%v\n", err)
		return
	}
	fmt.Printf("%#v\n", string(b))

	// 反序列化
	str := `{"name":"理想", "age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) //传指针是为了能在json.Unmarshal函数内部修改值
	fmt.Printf("%#v\n", p2)
}
