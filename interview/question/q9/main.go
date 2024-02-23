package main

import (
	"fmt"
	"sync"
)

// 交替打印数字和字母
// 问题描述
// 使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：
// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

var cA = make(chan struct{})
var cB = make(chan struct{})
var wg sync.WaitGroup
var s = "12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728"
var q = make(chan string, len(s))

func PrintA() {
	defer wg.Done()

	for {
		<-cA
		v, ok := <-q
		if !ok {
			close(cB)
			return
		}
		// 关闭channel
		fmt.Println("A:", v)
		cB <- struct{}{}
	}
}

func PrintB() {
	defer wg.Done()
	for {
		<-cB
		v, ok := <-q
		if !ok {
			close(cA)
			return
		}
		// 关闭channel
		fmt.Println("B:", v)
		cA <- struct{}{}
	}
}
func main() {
	for _, c := range s {
		q <- string(c)
	}
	close(q)

	wg.Add(2)
	go PrintA()
	go PrintB()
	cA <- struct{}{}
	wg.Wait()
}
