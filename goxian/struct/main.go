package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var p1 Person
	p1.Name = "zhujiajun"
	p1.Age = 18
	fmt.Println(p1)

	p2 := Person{"zhujiajun1", 19}
	fmt.Println(p2)

	var p3 *Person = new(Person)
	p3.Name = "zhujiajun3"
	p3.Age = 20
	fmt.Println(*p3)

	var p4 *Person = &Person{}
	p4.Name = "zhujiajun4"
	p4.Age = 22
	fmt.Println(*p4)
}
