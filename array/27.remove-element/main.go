package main

import "fmt"

//27. 移除元素
//https://leetcode-cn.com/problems/remove-element/submissions/
func removeElement(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}

func main() {
	fmt.Println(removeElement([]int{1, 2, 3, 4, 5, 3, 2, 1}, 3))
}
