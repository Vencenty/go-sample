package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type F struct {
	Base  int
	Value int
	Pos   int
}

func addTwoNumbers(l1, l2 *Node) *Node {
	f := &Node{
		Data: 0,
		Next: nil,
	}
	carry := 0
	current := f

	for l1.Next != nil || l2.Next != nil {
		x, y := 0, 0
		// 哨兵节点

		if l1 != nil {
			x = l1.Data
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Data
			l2 = l2.Next
		}

		sum := x + y + carry

		carry = sum / 10
		current.Next = &Node{
			Data: sum % 10,
			Next: nil,
		}
		current = current.Next
	}

	// 如果最后还有进位，需要添加一个节点
	if carry > 0 {
		current.Next = &Node{carry, nil}
	}

	return f.Next
}

func NewNode() *Node {
	return nil
}

func (n *Node) Insert(data int) *Node {
	node := &Node{
		Data: data,
		Next: n,
	}
	return node
}

// https://leetcode.cn/problems/add-two-numbers/description/
func main() {
	//node := &Node{}
	//node.Insert(2).Insert(4).Insert(3)

	a := 100
	b := 3

	fmt.Println(a / b)

	l1 := &Node{
		Data: 2,
		Next: &Node{
			Data: 4,
			Next: &Node{
				Data: 3,
				Next: nil,
			},
		},
	}
	l2 := &Node{
		Data: 5,
		Next: &Node{
			Data: 6,
			Next: &Node{
				Data: 4,
				Next: nil,
			},
		},
	}
	r := addTwoNumbers(l1, l2)
	fmt.Println(r)

}
