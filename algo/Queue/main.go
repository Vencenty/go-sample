package main

import "fmt"

type Queue struct {
	Elements []interface{}
}

func (q *Queue) Enqueue(elements ...interface{}) {
	q.Elements = append(q.Elements, elements...)
}

func (q *Queue) Dequeue() interface{} {
	if len(q.Elements) == 0 {
		return nil
	}
	element := q.Elements[0]
	q.Elements = q.Elements[1:]
	return element
}

func (q *Queue) List() {
	for _, v := range q.Elements {
		fmt.Println(v)
	}
}

func main() {
	q := new(Queue)
	q.Enqueue(1, 2, 3)
	q.Enqueue(4, 5, 6)
	q.Enqueue(7, 8, 9)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println("-------")
	q.List()
}
