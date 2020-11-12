package main

// 35. 搜索插入位置
// https://leetcode-cn.com/problems/search-insert-position/
// 输入: [1,3,5,6], 5
// 输出: 2
// 思路：暴力破解
func searchInsert(nums []int, target int) int {
	l := len(nums)
	for i := 0; i < l; i++ {
		if target <= nums[i] {
			return i
		}
	}
	return l
}

// 思路：二分法
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		middle := (right + left) / 2
		if nums[middle] >= target {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return left
}
