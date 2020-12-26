package main

import "math"

// 279. 完全平方数
// https://leetcode-cn.com/problems/perfect-squares/
func numSquares(n int) int {
	m := int(math.Sqrt(float64(n)))
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt64
	}
	for i := 1; i <= m; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}
	return dp[n]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
