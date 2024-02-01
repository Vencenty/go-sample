package main

type Node struct {
	Prev *Node
	Next *Node
	Data int
}

func (n *Node) Insert(value int) *Node {
	node := &Node{
		Prev: n,
		Next: nil,
		Data: value,
	}
	if n == nil {
		n.Next = node
	}
	return node
}

func main() {

}
