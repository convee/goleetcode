package main

// 704. 二分查找
// https://leetcode-cn.com/problems/binary-search/

// 总结一下一般实现的几个条件：
// 初始条件：left = 0, right = length-1
// 终止：left > right
// 向左查找：right = mid-1
// 向右查找：left = mid +1
func search(nums []int, target int) int {
	l, r, mid := 0, len(nums)-1, (len(nums)-1)/2
	for r >= l {
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
		mid = (l + r) / 2
	}
	return -1
}
