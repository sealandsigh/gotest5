package main

import (
	"fmt"
)

/*
你有50枚金币，需要分配给以下几个人，Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabetch.
分配规则如下:
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
b. 名字中每包含1个'o'或'O'分3枚金币
b. 名字中每包含1个'u'或'U'分4枚金币

写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 'dispatchCoin' 函数
*/

var (
	coins        = 50
	users        = []string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabetch"}
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin()
	fmt.Println("剩下:", left)
}

func dispatchCoin() (left int) {
	// 1 依次拿到每个人的名字
	// 2 查到一个人名根据分金币的规则去分金币
	// 2.1 每个人分的金币数应该保存到 distribution中
	// 2.2 还要记录下剩余的金币数量
	// 3 整个第二步执行完就能得到最终每个人分的金币数和剩余金币数

	for i := 0; i < len(users); i++ {
		for _, v := range users[i] {
			s := string(v)
			// fmt.Printf("%T\n", s)
			// if s == "e" || s == "E" {
			// 	fmt.Println("ok")
			// }
			switch s {
			case "e", "E":
				{
					if _, ok := distribution[users[i]]; !ok {
						distribution[users[i]] = 1
						coins -= 1
					} else {
						distribution[users[i]] += 1
						coins -= 1
					}
				}
			case "i", "I":
				{
					if _, ok := distribution[users[i]]; !ok {
						distribution[users[i]] = 2
						coins -= 2
					} else {
						distribution[users[i]] += 2
						coins -= 2
					}
				}
			case "o", "O":
				{
					if _, ok := distribution[users[i]]; !ok {
						distribution[users[i]] = 3
						coins -= 3
					} else {
						distribution[users[i]] += 3
						coins -= 3
					}
				}
			case "u", "U":
				{
					if _, ok := distribution[users[i]]; !ok {
						distribution[users[i]] = 4
						coins -= 4
					} else {
						distribution[users[i]] += 4
						coins -= 4
					}
				}

			}
		}
	}
	for k, v := range distribution {
		fmt.Printf("%s总共获得%d枚金币\n", k, v)
	}
	left = coins
	return
}
