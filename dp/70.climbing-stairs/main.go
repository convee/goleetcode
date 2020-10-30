package main

//70. 爬楼梯
//https://leetcode-cn.com/problems/climbing-stairs/
// 思路：dp[n]=dp[n-1]+dp[n-2]
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
