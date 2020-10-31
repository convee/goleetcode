package main

// 231. 2的幂
// 给定一个整数，编写一个函数来判断它是否是 2 的幂次方。
// https://leetcode-cn.com/problems/power-of-two/
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
