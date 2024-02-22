package main

import (
	"fmt"
	"sync"
)

var notifyA = make(chan struct{})
var notifyB = make(chan struct{})
var s = "ab"
var wg sync.WaitGroup
var q = make(chan string, len(s))

func PrintA() {
	defer wg.Done()

	for {
		<-notifyA
		v, ok := <-q
		if !ok {
			fmt.Println("A准备关闭")
			close(notifyB)
			return
		}
		fmt.Println("A:", v)
		notifyB <- struct{}{}
	}
}

func PrintB() {
	defer wg.Done()
	for {
		<-notifyB
		v, ok := <-q
		if !ok {
			fmt.Println("B准备关闭")
			close(notifyA)
			return
		}
		fmt.Println("B:", v)
		notifyA <- struct{}{}
	}
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
