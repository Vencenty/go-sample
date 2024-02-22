package main

import (
	"sort"
)

func main() {
	a1 := []int{1, 6}
	a2 := []int{3, 5}

	merged := make([]int, 0)

	merged = append(merged, a1...)
	merged = append(merged, a2...)
	sort.Ints(merged)
}
