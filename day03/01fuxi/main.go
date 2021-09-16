package main

import "fmt"

func main() {
	// 数组是值类型
	x := [3]int{1, 2, 3}
	fmt.Println(x)
	f1(x)
	fmt.Println(x)
	var s1 []int
	// s1 = make([]int, 1)
	// s1[0] = 100
	// fmt.Println(s1)
	s1 = append(s1, 1) // 自动初始化切片，这样的话上面就不用make初始化操作
	fmt.Println(s1)

}

func f1(a [3]int) {
	// Go语言中的函数传递的都是值 Ctrl+c Ctrl+v
	a[1] = 100

	// 指针
	// Go语言里面的指针只能读不能修改，不能修改指针变量指向的地址
	addr := "沙河"
	addrp := &addr
	fmt.Println(addrp)
	fmt.Printf("%T\n", *addrp)
	addrv := *addrp
	fmt.Println(addrv)

	// map
	var m1 map[string]int
	fmt.Println(m1)
	m1 = make(map[string]int, 3)
	m1["理想"] = 100
	fmt.Println(m1)
	fmt.Println(m1["ji"]) //如果key不存在返回对应类型的零值
	// 如果返回值是布尔型，我们通常用ok去接收它
	souce, ok := m1["ji"]
	if !ok {
		fmt.Println("没有姬无命这个人")
	} else {
		fmt.Println("姬无命的分数为", souce)
	}
	delete(m1, "lixiang") // 删除的key不存在什么都不干
	delete(m1, "理想")
	fmt.Println(m1)
	fmt.Println(m1 == nil) // 已经开辟了内存
}
