package main

//70. 爬楼梯
//https://leetcode-cn.com/problems/climbing-stairs/
// 思路：dp[n]=dp[n-1]+dp[n-2]
// 上 1 阶台阶：有1种方式。
// 上 2 阶台阶：有1+1和2两种方式。
// 上 3 阶台阶：到达第3阶的方法总数就是到第1阶和第2阶的方法数之和。
// 上 n 阶台阶，到达第n阶的方法总数就是到第 (n-1) 阶和第 (n-2) 阶的方法数之和。
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
