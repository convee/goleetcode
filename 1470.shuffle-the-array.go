package main

import (
	"fmt"
)

//重新排列数组
//https://leetcode-cn.com/problems/shuffle-the-array/
func shuffle(nums []int, n int) []int {
	res := make([]int, 2*n)
	for i := 0; i < n; i++ {
		res[2*i] = nums[i]
		res[2*i+1] = nums[n+i]
	}
	return res
}

func main() {
	nums := []int{2, 5, 1, 3, 4, 7}
	n := 3
	fmt.Println(shuffle(nums, n))
}
