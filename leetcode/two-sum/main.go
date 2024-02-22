package main

import "fmt"

//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//你可以按任意顺序返回答案。
//示例 method1：
//输入：nums = [2,7,11,15], target = 9
//输出：[0,method1]
//解释：因为 nums[0] + nums[method1] == 9 ，返回 [0, method1] 。
//示例 2：
//输入：nums = [3,2,4], target = 6
//输出：[method1,2]
//示例 3：
//输入：nums = [3,3], target = 6
//输出：[0,method1]
//
//func TwoSum(nums []int, target int) []int {
//	result := make(map[int]int)
//	for i, num := range nums {
//		diff := target - num
//		if j, ok := result[diff]; ok {
//			return []int{i, j}
//		}
//		result[num] = i
//	}
//	return nil
//}

func TwoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, num := range nums {
		if j, found := numsMap[target-num]; found {
			return []int{i, j}
		}
		numsMap[num] = i
	}
	return nil
}

func main() {
	target := 18
	arr := []int{2, 7, 11, 15}
	fmt.Println(TwoSum(arr, target))
}
