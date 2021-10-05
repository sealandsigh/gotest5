package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// ini配置文件解析器

// MysqlConfig MySQL配置结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

// RedisConfig ...
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

// Config ...
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(filename string, data interface{}) (err error) {
	// 0 参数的校验
	// 0.1 传进来的data参数必须是指针类型(因为需要在函数中对其赋值)
	t := reflect.TypeOf(data)
	// t.Kind() 可以查看数据类型，这里的类型是大类的类型，比如你自己可以定义cat dog但都是属于struct类型这个大类
	fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr {
		err = errors.New("data should be a pointer") // 格式化输出之后返回一个error类型,创建一个错误
		return
	}
	// 0.2 传进来的data参数必须是结构体类型指针(因为配置文件中各种键值才需要赋值给结构体的字段)
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data should be a struct") // 格式化输出之后返回一个error类型,创建一个错误
		return
	}
	// 1 读文件得到字节类型数据
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	// string(b) // 将文件内容转换成字符串
	lineSlice := strings.Split(string(b), "\n")
	fmt.Println(lineSlice)
	// 2 一行一行得读数据
	var structName string
	for idx, line := range lineSlice {
		// 去掉字符串首尾空格
		line = strings.TrimSpace(line)
		// 如果是空行就跳过
		if len(line) == 0 {
			continue
		}
		// 2.1 如果是注释就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// 2.2 如果是[开头就表示是节(section)]
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 把这一行首尾[]去掉，去到中间的内容把首尾的空格去掉拿到内容
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 根据字符串的sectionName去data里面根据反射找到对应结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					// 说明找到了对应的嵌套结构体，把字段名记下来
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}

		} else {
			// 2.3 如果不是[开头就是=分割的键值对]
			// 1 以等号分割这一行，等号左边是key,等号右边是value
			if !strings.Contains(line, "=") || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			// 2 根据strucName去data里面把对应的嵌套结构体取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) // 拿到嵌套结构体信息
			sType := sValue.Type()                     // 拿到嵌套结构体的类型信息

			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s字段应该是一个结构体", structName)
				return
			}
			var fieldName string
			var fileType reflect.StructField
			// 3 遍历嵌套结构体的每一个字段，判断tag是不是等于key
			for i := 0; i < sValue.NumField(); i++ {
				field := sType.Field(i) // tag信息是存储在类型信息中
				fileType = field
				if field.Tag.Get("ini") == key {
					// 找到对应的字段
					fieldName = field.Name
					break
				}
			}
			// 4 如果key = tag, 给这个字段赋值
			// 4.1 如果key = tag，给这个字段赋值
			if len(fieldName) == 0 {
				// 在结构体中找不到对应的字段
				continue
			}
			fileObj := sValue.FieldByName(fieldName)
			// 4.2 对其赋值
			fmt.Println(fieldName, fileType.Type.Kind())
			switch fileType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				var valueFloat float64
				valueFloat, err = strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetFloat(valueFloat)
			}
		}
	}
	return
}

func main() {
	var cfg Config
	// var x = new(int)
	err := loadIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini faild,err:%v", err)
	}
	fmt.Println(cfg)
}
