package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func NewList(value int) *Node {
	return &Node{Data: value, Next: nil}
}

func (n Node) Insert(value int) *Node {
	node := &Node{
		Data: value,
		Next: nil,
	}
	return node
}

func main() {
	linkList := NewList(1)
	linkList = linkList.Insert(2)
	linkList = linkList.Insert(3)
	linkList = linkList.Insert(4)

	for cursor := linkList; cursor != nil; cursor = cursor.Next {
		fmt.Println(cursor.Data)
	}

}
