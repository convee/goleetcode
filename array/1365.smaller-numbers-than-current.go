package main

//1365. 有多少小于当前数字的数字
//https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number/
func smallerNumbersThanCurrent(nums []int) []int {
	var bucket [101]int
	res := make([]int, len(nums))
	for _, value := range nums {
		bucket[value]++
	}
	for i := 1; i < 101; i++ {
		bucket[i] += bucket[i-1]
	}
	for i, v := range nums {
		if v != 0 {
			res[i] = bucket[nums[i]-1]
		}
	}
	return res
}
