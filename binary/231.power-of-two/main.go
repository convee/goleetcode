package main

// 231. 2的幂
// 给定一个整数，编写一个函数来判断它是否是 2 的幂次方。
// https://leetcode-cn.com/problems/power-of-two/
// 可以看到，对于N为2的幂的数，都有 N&(N-1)=0 ，所以这就是我们的判断条件。
// 例如 2^3 = 8,
// 8 的二进制 1 0 0 0
// 7 的二进制 0 1 1 1
// 7 & 8 = 0
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
