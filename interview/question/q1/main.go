package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var q = make(chan int)
var wg sync.WaitGroup

func main() {
	method2()
}

// 用wg实现
func method1() {
	go func() {
		for i := 0; i < 5; i++ {
			q <- rand.Intn(100)
		}
		close(q)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range q {
			fmt.Println(v)
		}
	}()
	wg.Wait()
}

// 不用wg，用<-done实现
func method2() {
	done := make(chan struct{})

	go func() {
		for i := 0; i < 5; i++ {
			q <- rand.Intn(100)
		}
		close(q)
	}()

	go func() {
		for v := range q {
			fmt.Println(v)
		}
		done <- struct{}{}
	}()

	<-done
}
