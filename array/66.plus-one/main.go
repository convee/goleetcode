package main

//66. 加一
//https://leetcode-cn.com/problems/plus-one/
func plusOne(digits []int) []int {
	l := len(digits)
	for i := l - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
	}
	return append([]int{1}, digits...)
}
