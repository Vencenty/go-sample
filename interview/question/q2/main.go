package main

import (
	"fmt"
	"strings"
)

// 判断字符串中字符是否全都不同
// 问题描述
// 请实现一个算法，确定一个字符串的所有字符【是否全都不同】。这里我们要求【不允许使用额外的存储结构】。 给定一个string，请返回一个bool值,true代表所有字符全都不同，false代表存在相同的字符。 保证字符串中的字符为【ASCII字符】。字符串的长度小于等于【3000】。
func main() {
	//r := method2("1233")
	r := method3("123asdf")
	fmt.Println(r)
}

// 方法3
func method3(s string) bool {
	for _, v := range s {
		if v > 127 {
			return false
		}
		if strings.Index(s, string(v)) != strings.LastIndex(s, string(v)) {
			return false
		}

	}
	return true
}

// 方法2
func method2(s string) bool {
	for k, v := range s {
		if v > 127 {
			return false
		}
		if strings.Index(s, string(v)) != k {
			return false
		}
	}
	return true
}

// true代表无重复，false代表有重复
func method1(s string) bool {
	setMap := map[string]struct{}{}

	for _, c := range s {
		_, ok := setMap[string(c)]
		if ok {
			return false
		}
		setMap[string(c)] = struct{}{}
	}
	return true
}
