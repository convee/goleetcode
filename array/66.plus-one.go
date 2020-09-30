package main

//66. 加一
//输入: [1,2,3]
//输出: [1,2,4]
//解释: 输入数组表示数字 123。
//https://leetcode-cn.com/problems/plus-one/
func plusOne(digits []int) []int {
	l := len(digits)

	if l == 0 {
		return []int{1}
	}

	for i := l - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}
	return append([]int{1}, digits...)
}

func plusOne2(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	nums := make([]int, len(digits)+1)
	nums[0] = 1
	return nums
}
