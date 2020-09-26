package offer

//剑指 Offer 10- I. 斐波那契数列
//https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/
// 写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项。斐波那契数列的定义如下：

// F(0) = 0,   F(1) = 1
// F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
// 斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。

// 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

// 解题思路
// 先判断0和1的情况
// 循环，从第二次开始，v2赋给v1，新的v2 = (v1+v2) % 1000000007
// 最后判断等于1000000008，返回1

func fib(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 1, 1
	for i := 2; i < n; i++ {
		a, b = b, (a+b)%1000000007
	}
	if b == 1000000008 {
		return 1
	}
	return b
}
