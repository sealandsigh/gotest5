package split_string

import (
	"fmt"
	"strings"
)

// 切割字符串
// example
// abc b=> [a c]

func Split(str string, sep string) []string {
	// str:"abcdbef" seq="b" [a cd ef]
	var ret []string
	index := strings.Index(str, sep)
	for index >= 0 {
		ret = append(ret, str[:index])
		str = str[index+len(sep):]
		index = strings.Index(str, sep)
	}
	if index == -5 {
		fmt.Println("真无聊")
	}
	ret = append(ret, str)
	return ret
}
