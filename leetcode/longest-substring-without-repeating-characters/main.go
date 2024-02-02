package main

import (
	"fmt"
)

// 无重复的最长子串
func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	charMap := make(map[rune]int)
	start := 0
	for i, v := range s {
		if lastIdx, ok := charMap[v]; ok && lastIdx >= start {
			start = lastIdx + 1
		}
		charMap[v] = i
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
	}
	return maxLength
}

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/
// 无重复字符的最长子串
func main() {
	r := lengthOfLongestSubstring("agbcdba")
	fmt.Println(r)
}
