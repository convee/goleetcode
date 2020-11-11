package main

// 121. 买卖股票的最佳时机
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
// 思路：动态规划
// 前i天的最大收益 = max(前i-1天的最大收益，第i天的价格-前i-1天中的最小价格)
// 记录【今天之前买入的最小值】
// 计算【今天之前最小值买入，今天卖出的获利】，也即【今天卖出的最大获利】
// 比较【每天的最大获利】，取最大值即可
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	mi := prices[0]
	ma := 0
	for i := 1; i < len(prices); i++ {
		ma = max(ma, prices[i]-mi)
		mi = min(mi, prices[i])
	}
	return ma
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
