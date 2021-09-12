package main

import "fmt"

func main() {
	// s := "hello沙河"
	// for i := 0; i < len(s); i++ { //byte
	// 	fmt.Printf("%v(%c) ", s[i], s[i])
	// }
	// fmt.Println()
	// for _, r := range s { //rune
	// 	fmt.Printf("%v(%c) ", r, r)
	// }
	// fmt.Println()

	// "Hello" => 'H' 'e' 'l' 'l' 'o'
	// 字符串修改
	s2 := "白萝卜"      // => '白' '萝' '卜'
	s3 := []rune(s2) // 把字符串强制改成一个rune分片
	s3[0] = '红'
	fmt.Println(string(s3))

	c1 := "红"
	c2 := '红'
	fmt.Printf("c1:%T c2:%T\n", c1, c2)
	c3 := "H"
	c4 := 'H'
	fmt.Printf("c3:%T c4:%T\n", c3, c4)

	//类型转换
	n := 10
	var f float64
	f = float64(n)
	fmt.Println(f)
}
