package main

import (
	"fmt"
	"sync"
)

var notifyA = make(chan struct{})
var notifyB = make(chan struct{})
var s = "abc"
var wg sync.WaitGroup
var q = make(chan string, len(s))
var done = make(chan struct{})

func PrintA() {
	defer func() {
		r := recover()
		fmt.Println("recoverA", r)
	}()
	defer wg.Done()

	<-notifyA
	for v := range q {
		fmt.Println("A:", v)
		notifyB <- struct{}{}
	}
	close(notifyB)
}

func PrintB() {
	defer func() {
		r := recover()
		fmt.Println("recoverB", r)
	}()
	defer wg.Done()

	for v := range q {
		<-notifyB
		fmt.Println("B:", v)
		notifyA <- struct{}{}
	}
	close(notifyA)

}

func main() {
	wg.Add(2)
	for _, c := range s {
		q <- string(c)
	}
	close(q)
	go PrintA()
	go PrintB()
	notifyA <- struct{}{}
	wg.Wait()
}
