package main

// 53. 最大子序和
// https://leetcode-cn.com/problems/maximum-subarray/
// 输入: [-2,1,-3,4,-1,2,1,-5,4]
// 输出: 6
// 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
// 思路：dp[i]=max(nums[i], dp[i−1]+nums[i])
func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	result := nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		result = max(dp[i], result)
	}
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
