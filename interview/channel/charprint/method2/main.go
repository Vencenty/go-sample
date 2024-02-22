package main

import (
	"fmt"
	"sync"
)

var notifyA = make(chan struct{})
var notifyB = make(chan struct{})
var s = "abcdefg"
var wg sync.WaitGroup
var q = make(chan string, len(s))
var done = make(chan struct{})

func PrintA() {
	defer wg.Done()

	for {
		select {
		case <-notifyA:
			v, ok := <-q
			if !ok {
				done <- struct{}{}
				return
			}
			fmt.Println("A:", v)
			notifyB <- struct{}{}
		case <-done:
			close(notifyB)
			return
		}
	}
}

func PrintB() {
	defer wg.Done()
	for {
		select {
		case <-notifyB:
			v, ok := <-q
			if !ok {
				done <- struct{}{}
				return
			}
			fmt.Println("B:", v)
			notifyA <- struct{}{}
		case <-done:
			close(notifyA)
			return
		}
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
