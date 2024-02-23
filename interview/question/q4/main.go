package main

import (
	"fmt"
)

func main() {

	str := "1234"
	s := reserveString(str)
	fmt.Println(s)

}

// 反转一个字符串
func reserveString(str string) string {
	l := len(str)    //5
	s := []rune(str) // 这里面有坑，string不能直接反转
	for i := 0; i < l/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
	return string(s)

}
