package main

//278. 第一个错误的版本
//https://leetcode-cn.com/problems/first-bad-version/
func isBadVersion(version int) bool

// 思路：二分查找
// 不能写 mid := (left + right) / 2; 要写 mid := left + (right - left) / 2;
// 返回 left 或者 right 都行，因为终止条件是 left == right
func firstBadVersion(n int) int {
	left := 1
	right := n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left

}
