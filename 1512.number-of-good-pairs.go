package main

import "fmt"

//https://leetcode-cn.com/problems/number-of-good-pairs/

func numIdenticalPairs(nums []int) int {
	l := len(nums)
	if l <= 0 {
		return 0
	}
	if l == 1 {
		return 1
	}
	n := 0
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i] == nums[j] {
				n++
			}
		}
	}
	return n
}

func main() {
	nums := []int{1, 2, 3, 1, 1, 3}
	fmt.Println(numIdenticalPairs(nums))
}
