package main

import (
	"fmt"
	"strings"
)

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	binarySort(arr, 1)

	s := "hello"
	fmt.Println(strings.Index(s, "l"), strings.LastIndex(s, "l"))
}

func binarySort(arr []int, target int) {
	l := len(arr)

	for i := 0; i < l/2; i++ {
	}
}
