package main

import "sync"

func Worker() {

}

func Generator() {
	for {
		select {
		case <-job:

		}
	}
}

var job = make(chan string, 10)
var wg sync.WaitGroup

func main() {

	for i := 0; i < 5; i++ {
		go Generator()
	}

	for i := 0; i < 3; i++ {
		wg.Add(2)
		go Worker()
		wg.Wait()
	}
}
