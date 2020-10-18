package main

import "fmt"

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 *
 * https://leetcode-cn.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (44.40%)
 * Total Accepted:    239.4K
 * Total Submissions: 539.3K
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
 *
 * 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
 *
 * 示例:
 *
 * 给定 nums = [2, 7, 11, 15], target = 9
 *
 * 因为 nums[0] + nums[1] = 2 + 7 = 9
 * 所以返回 [0, 1]
 *
 *
 */

func twoSum(nums []int, target int) []int {
	if len(nums) <= 0 {
		return []int{}
	}
	m := make(map[int]int)
	for index, num := range nums {
		if _, ok := m[target-num]; ok {
			return []int{index, m[target-num]}
		}
		m[num] = index
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	l := len(nums)
	result := []int{}
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i)
				result = append(result, j)
				return result
			}
		}
	}
	return result
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Print(twoSum(nums, target))
}
