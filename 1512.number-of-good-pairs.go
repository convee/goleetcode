package main

import "fmt"

//1512. 好数对的数目
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

func numIdenticalPairs2(nums []int) int {
	res := 0

	var m map[int]int = make(map[int]int, 10)
	for i := 0; i < len(nums); i++ {
		if v, ok := m[nums[i]]; ok {
			res += v
			m[nums[i]] = v + 1
		} else {
			m[nums[i]] = 1
		}

	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 1, 1, 3}
	fmt.Println(numIdenticalPairs2(nums))
}
