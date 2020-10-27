package main

import (
	"fmt"
	"sort"
)

//561. 数组拆分 I
//先将数组排序，然后把从a1到an数组下标为奇数的数都加起来
//https://leetcode-cn.com/problems/array-partition-i/
func arrayPairSum(nums []int) int {
	l := len(nums)
	sum := 0
	sort.Ints(nums)
	for i := 0; i < l; i++ {
		if i%2 == 0 {
			sum += nums[i]
		}
	}
	return sum
}

func main() {
	nums := []int{2, 3, 4, 1}
	fmt.Println(arrayPairSum(nums))
}
