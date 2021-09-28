package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind() // 值的类型种类
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整形的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64,value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Int()从反射中获取整形的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64,value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Int()从反射中获取整形的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64,value is %f\n", float64(v.Float()))
	}
}

func main() {
	var a float32 = 3.14
	reflectType(a) // type:float32
	var b int64 = 100
	reflectType(b) // type:int64

	// ValueOf
	reflectValue(a)
}
