package main

import "fmt"

//724. 寻找数组的中心索引
//https://leetcode-cn.com/problems/find-pivot-index/
func pivotIndex(nums []int) int {
	l := len(nums)
	if l < 3 {
		return -1
	}
	for i := 0; i < l; i++ {
		jn := 0
		for j := 0; j < i; j++ {
			jn += nums[j]
		}
		kn := 0
		for k := i + 1; k < l; k++ {
			kn += nums[k]
		}
		if jn == kn {
			return i
		}
	}
	return -1
}

func main() {
	nums := []int{-1, -1, -1, 0, 1, 1}

	fmt.Println(pivotIndex(nums))
}
