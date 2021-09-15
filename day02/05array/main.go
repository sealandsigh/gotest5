package main

import "fmt"

// 数组

// 存放元素的容器
// 必须指定存放的元素的类型和容量 (长度)
// 数组的长度是数组类型的一部分

func main() {
	var a1 [3]bool // [true, false, true]
	var a2 [4]bool // [true, false, true, true]

	fmt.Printf("a1:%T, a2:%T\n", a1, a2)

	// 数组的初始化
	// 如果不初始化，默认元素都是零值 (布尔值:false 整形和浮点数 0 字符串 "")
	fmt.Println(a1, a2)
	// 1 初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	// 2 初始化方式2
	// a10 := [9]int{0, 1, 2, 3, 4, 4, 5, 6, 7}
	// 根据初始值自动推断数组的长度是多少
	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a10)
	// 3 初始化方式3
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)

	// 数组的遍历
	city := [...]string{"北京", "上海", "深圳"}
	//1. 根据索引遍历
	for i := 0; i < len(city); i++ {
		fmt.Println(city[i])
	}
	// for range
	for i, v := range city {
		fmt.Println(i, v)
	}

	// 多维数组
	// [[1 2], [3 4], [5 6]]
	var a11 [3][2]int = [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	// a11 = [3][2]int{
	// 	{1, 2},
	// 	{3, 4},
	// 	{5, 6},
	// }
	fmt.Println(a11)

	// 多维数组的遍历
	for _, v1 := range a11 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	// 数组是值类型, ctrl+c ctrl+v 把文件从文件夹a拷贝到文件夹b
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)

	// array数组练习题
	// 1. 求数组[1, 3, 5, 7, 8]所有元素的和
	// 找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7]中找出和为8的两个元素的下标分别为(0, 3)和(1, 2)

	c1 := [...]int{1, 3, 5, 7, 8}
	total := 0
	for _, v := range c1 {
		total += v
	}
	fmt.Println(total)

	// 找出和为8的两个元素的下标分别为(0,3) 和 (1,2)
	// 定义两个for循环，外层的从第一个开始遍历
	// 内层的for循环从外层后面的那个开始找
	// 它们两个数的和为8
	for i := 0; i < len(c1); i++ {
		for j := i + 1; j < len(c1); j++ {
			if c1[i]+c1[j] == 8 {
				fmt.Printf("(%d, %d)\n", i, j)
			}
		}
	}
}
