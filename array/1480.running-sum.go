package main

import "fmt"

// 一位数组的动态和
// 输入 [1,2,3,4]
// 输出 [1,3,6,10
func runningSum(nums []int) []int {
	if len(nums) <= 0 {
		return []int{}
	}
	for i := range nums {
		if i != 0 {
			nums[i] = nums[i-1] + nums[i]
		}
	}
	return nums
}
func runningSum2(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}
	return nums
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(runningSum2(nums))
}
