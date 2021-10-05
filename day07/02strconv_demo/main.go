package main

import (
	"fmt"
	"strconv"
)

// strconv

func main() {
	// 从字符串解析出对应的数字
	str := "10000"
	// 这里的64是转换的位数，但是这里返回的永远是int64，但是这里是按照你定义的位数转换，
	// 只不过需要后续再自行转换，只不过就没有丢失位数的风险了，因为前面已经转换了，虽说返回值永远是int64
	ret1, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Printf("parseint failed,err:%v", err)
		return
	}
	fmt.Println(ret1)

	//
	retIint, _ := strconv.Atoi(str)
	fmt.Println(retIint)

	// 将字符串中解析出bool值
	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n", boolValue, boolValue)

	// 从字符串中解析出浮点数
	floatStr := "1.234"
	floatValue, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("%#v %T\n", floatValue, floatValue)

	// 把数字转换成字符串类型
	i := int32(97)

	ret2 := fmt.Sprintf("%d", i) // "97"
	fmt.Printf("%#v\n", ret2)
	ret3 := strconv.Itoa(int(i))
	fmt.Printf("%#v\n", ret3)
}
