package main

import (
	"fmt"
	"time"
)

// 时间

func main() {
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	// 时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	// time.Unix()
	ret := time.Unix(1632638236, 0)
	fmt.Println(ret)
	fmt.Println(ret.Year())
	fmt.Println(ret.Day())
	// 时间间隔
	fmt.Println(time.Second)
	// now + 1小时
	fmt.Println(now.Add(24 * time.Hour))
	// sub 求两个时间差值
	timeObj1, err := time.Parse("2006-01-02", "2021-09-27")
	if err != nil {
		fmt.Printf("parse time failed err:%v\n", err)
		return
	}
	d := now.Sub(timeObj1)
	fmt.Println(d)
	// equal 判断两个时间的比较
	// before 判断哪个时间点在时间之前
	// after 判断哪个时间在时间之后
	// 定时器
	// timer := time.Tick(time.Second)
	// for t := range timer {
	// 	fmt.Println(t)
	// }

	// 格式化时间 把语言中时间对象 转换成字符串类型的的时间
	// 2019-08-03
	fmt.Println(now.Format("2006-01-02"))
	// 2019/02/03 11:55:02
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	// 2019/02/03 14:54:02PM
	fmt.Println(now.Format("2006/01/02 03:04:05 PM"))
	// 2019/02/03 14:54:02.342
	fmt.Println(now.Format("2006/01/02 15:04:05.000"))
	// 按照对应的格式解析字符串类型的时间
	timeObj, err := time.Parse("2006-01-02", "2000-08-03")
	if err != nil {
		fmt.Printf("parse time failed err:%v\n", err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())
	// sleep
	// time.Sleep(5 * time.Second)
}
