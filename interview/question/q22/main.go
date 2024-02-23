package main

import (
	"fmt"
	"reflect"
	"sync"
)

func main() {
	var m sync.Map
	m.Store("address", map[string]string{
		"province": "Shandong",
		"city":     "weifang",
	})

	r, ok := m.Load("address")
	if !ok {
		fmt.Println("找不到对应值")
		return
	}

	fmt.Println(reflect.TypeOf(r))
	fmt.Println(r.(map[string]string)["city"])
}
