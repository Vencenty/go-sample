package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
// 例如，121 是回文，而 123 不是。

func isPalindrome(x int) bool {
	return true
}

// isPalindromeByStringReverse 判断是否是回文，通过字符串来做
func isPalindromeByStringReverse(x int) bool {
	originStr := strconv.Itoa(x)
	strMap := make([]string, len(originStr))
	// 反转字符串
	strLen := len(originStr) - 1
	for i, v := range originStr {
		index := strLen - i
		strMap[index] = string(v)
	}
	reverseStr := strings.Join(strMap, "")
	return reverseStr == originStr
}

func main() {
	num := 1234
	//r := isPalindromeByString(num)
	f := isPalindromeByStringReverse(num)
	fmt.Println(f)
	//fmt.Println(r)
}
