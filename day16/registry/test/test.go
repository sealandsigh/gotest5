package main

import "fmt"

func main() {
	i := 0
	f(&i)
	fmt.Println(i)
}

func f(count *int) {
	fmt.Println(*count)
	(*count)++
}
