package main

import "fmt"

type ListMode struct {
	Val  int
	Next *ListMode
}

func reverseList(head *ListMode) *ListMode {
	var pre *ListMode
	cur := head
	for cur != nil {
		tmp := cur.Next // 把下一个节点缓存起来
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

func main() {
	head := &ListMode{
		Val: 1,
		Next: &ListMode{
			Val: 2,
			Next: &ListMode{
				Val: 3,
				Next: &ListMode{
					Val: 4,
					Next: &ListMode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	fmt.Printf("%#v\n", head)
	ret := reverseList(head)
	fmt.Printf("%#v\n", ret)
	for ret != nil {
		fmt.Print(ret.Val, "->")
		ret = ret.Next
	}
}
