package main

import "fmt"

//1464. 数组中两元素的最大乘积
//https://leetcode-cn.com/problems/maximum-product-of-two-elements-in-an-array/
func maxProduct(nums []int) int {
	l := 0
	xl := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > l {
			l = nums[i]
		}
		if l > xl {
			l, xl = xl, l
		}
	}
	return (l - 1) * (xl - 1)
}

func main() {
	nums := []int{3, 4, 5, 2}
	fmt.Println(maxProduct(nums))
}
