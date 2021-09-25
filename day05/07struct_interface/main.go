package main

import "fmt"

// 同一个结构体可以实现多个接口

// type animal interface {
// 	mover
// 	eater
// }

type mover interface {
	move()
}

type eater interface {
	eat(string)
}

type cat struct {
	name string
	feet int8
}

// cat 实现了mover接口
func (c *cat) move() {
	fmt.Println("走猫步。。")
}

// cat 实现了eater接口
func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s鱼\n", food)
}

func main() {
	// var a1 animal
	var m1 mover
	var e1 eater

	c1 := cat{"tome", 4}
	m1 = &c1
	e1 = &c1
	m1.move()
	e1.eat("hotdog")
}
