package main

import "fmt"

// 137. 只出现一次的数字 II
// https://leetcode-cn.com/problems/single-number-ii/
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。
func singleNumber2(nums []int) int {
	// 统计每位1的个数
	var result int
	for i := 0; i < 64; i++ {
		sum := 0
		for j := 0; j < len(nums); j++ {
			// 统计1的个数
			sum += (nums[j] >> i) & 1
		}
		// 还原位00^10=10 或者用| 也可以
		result ^= (sum % 3) << i
	}
	return result
}

func singleNumber(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}

//错误用法示例
func main() {
	nums := []int{1, 2, 2, 2, 3, 3, 3}
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for _, v := range nums {
		fmt.Println(m[v])
	}
}
