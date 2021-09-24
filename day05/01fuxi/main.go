package main

import (
	"encoding/json"
	"fmt"
)

// 复习结构体
// 结构体和匿名结构体

type tmp struct {
	x int
	y int
}

type person struct {
	name string
	age  int
}

// func sum(x int, y int) (ret int) {
// 	ret = x + y
// 	return ret
// }

// 这个例子阐述的是 person这个返回值其实是个未命名的返回值，也可以命名，就是麻烦，所以用未命名形式使用
// func newPerson(name string, age int) (p person) {
// 	p =  person{
// 		name: name,
// 		age:  age,
// 	}
// 	return p
// }

func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

// 方法
// 接收者使用对应类型首字母的小写
// 指定了了接收者之后，只有接收者这个类型的变量才能调用这个方法
func (p person) dream(str string) {
	fmt.Printf("%s的梦想是把%s学好\n", p.name, str)
}

// func (p person) guonian() {
// 	p.age++ //此处的p是下面c1的副本 改的是副本
// }

// 指针，内存地址，实际指向的是person类型
// 指针接收者
// 1 需要修改结构体变量的值时要使用指针接收者
// 2 结构体本身比较大, 拷贝的内存开销比较大时也要使用指针接收者
// 3 保持一致性 如果有一个方法使用了指针接收者，剩下的方法统一也使用指针接收者
func (p *person) guonian() {
	p.age++ //此处的p是直接可以更改年龄的，因为直接改的是原来的值
}

func main() {
	var b = tmp{
		10,
		20,
	}
	fmt.Println(b)

	var a = struct {
		x int
		y int
	}{10, 20}
	fmt.Println(a)

	var c = newPerson("周林", 18)
	fmt.Println(c)
	c.dream("python")
	fmt.Println(c.age)
	c.guonian()
	// (&c).guonian()  // 实际应该这么写，但是go语法糖可以直接写c调用
	fmt.Println(c.age)

	// 结构体嵌套
	type addr struct {
		province string
		city     string
	}
	type student struct {
		name    string
		address addr //嵌套别的结构体, 去掉address就叫匿名嵌套,匿名嵌套实际上就是使用类型当名称
	}

	d := student{
		name: "周林",
		address: addr{
			"朝阳",
			"北京",
		},
	}
	fmt.Println(d)

	type point struct {
		X int `json:"x"`
		Y int `json:"y"`
	}

	p1 := point{100, 200}
	b1, err := json.Marshal(p1)
	// 如果出错了
	if err != nil {
		fmt.Println("序列化出错了")
	} else {
		fmt.Println(string(b1))
	}

	// 反序列化 由字符串 -》 go中的结构体变量
	str1 := `{"x":10, "y":20}`
	var po2 point // 造一个结构体变量，准备接收反序列化的值
	err = json.Unmarshal([]byte(str1), &po2)
	if err != nil {
		fmt.Printf("unmarshal failed, err:%v\n", err)
	} else {
		fmt.Println(po2)
	}

}
