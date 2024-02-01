package main

import "fmt"

type Stack struct {
	Elements []interface{}
}

func (s *Stack) Add(element ...interface{}) {
	s.Elements = append(s.Elements, element...)
}

func (s *Stack) Pop() interface{} {
	if len(s.Elements) == 0 {
		return nil
	}
	// 最后一个元素的索引
	lastIndex := len(s.Elements) - 1
	button := s.Elements[lastIndex]
	s.Elements = s.Elements[:lastIndex]
	return button
}

func (s *Stack) List() {
	for _, v := range s.Elements {
		fmt.Println(v)
	}
}

func main() {
	s := &Stack{}
	s.Add(1, 2, 3)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
